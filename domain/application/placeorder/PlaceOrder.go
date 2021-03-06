package placeorder

import (
	"clean-arch-go/domain/entity"
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/gateway"
	"clean-arch-go/domain/repository"
	"clean-arch-go/domain/service"
	"errors"
	"fmt"
)

var ErrItemNotFound = errors.New("item not found")

type PlaceOrder struct {
	zipcodeClient     gateway.ZipcodeClient
	couponRepository  repository.CouponRepository
	itemRepository    repository.ItemRepository
	orderRepository   repository.OrderRepository
	freightCalculator service.FreightCalculator
}

func NewPlaceOrder(zipcodeClient gateway.ZipcodeClient, freightCalculator service.FreightCalculator, repositoryFactory factory.RepositoryFactory) PlaceOrder {
	return PlaceOrder{
		zipcodeClient:     zipcodeClient,
		couponRepository:  repositoryFactory.NewCouponRepository(),
		itemRepository:    repositoryFactory.NewItemRepository(),
		orderRepository:   repositoryFactory.NewOrderRepository(),
		freightCalculator: freightCalculator,
	}
}

func (uc PlaceOrder) Execute(input InputPlaceOrder) (OutputPlaceOrder, error) {
	sequence := uc.orderRepository.Count() + 1
	order, err := entity.NewOrder(input.Document, input.IssueDate, sequence)
	if err != nil {
		return OutputPlaceOrder{}, fmt.Errorf("could not create order: %w", err)
	}
	distance := uc.zipcodeClient.Distance(input.ZipcodeOrigin, input.ZipcodeDestiny)
	for _, itemInput := range input.Items {
		item, err := uc.itemRepository.GetById(itemInput.ItemID)
		if err != nil {
			return OutputPlaceOrder{}, err
		}
		if (item == entity.Item{}) {
			return OutputPlaceOrder{}, ErrItemNotFound
		}
		order.AddItem(itemInput.ItemID, item.Price, itemInput.Quantity)
		order.ShippingCost += uc.freightCalculator.Calculator(distance, item) * itemInput.Quantity
	}
	if input.CouponCode != "" {
		coupon, err := uc.couponRepository.FindByCode(input.CouponCode)
		if err != nil {
			return OutputPlaceOrder{}, err
		}
		order.AddCoupon(coupon)
	}
	total := order.GetTotal()
	err = uc.orderRepository.Save(order)
	if err != nil {
		return OutputPlaceOrder{}, err
	}
	return OutputPlaceOrder{Code: order.Code.Value, ShippingCost: order.ShippingCost, Total: total}, nil
}
