package application

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

func (uc PlaceOrder) Execute(input PlaceOrderInput) (PlaceOrderOutput, error) {
	sequence := uc.orderRepository.Count() + 1
	order, err := entity.NewOrder(input.Document, input.IssueDate, sequence)
	if err != nil {
		return PlaceOrderOutput{}, fmt.Errorf("could not create order: %w", err)
	}
	distance := uc.zipcodeClient.Distance(input.ZipcodeOrigin, input.ZipcodeDestiny)
	for _, itemInput := range input.Items {
		item, err := uc.itemRepository.GetById(itemInput.itemID)
		if err != nil {
			return PlaceOrderOutput{}, err
		}
		if (item == entity.Item{}) {
			return PlaceOrderOutput{}, ErrItemNotFound
		}
		order.AddItem(itemInput.itemID, item.Price, itemInput.quantity)
		order.ShippingCost += uc.freightCalculator.Calculator(distance, item) * itemInput.quantity
	}
	if input.CouponCode != "" {
		coupon, err := uc.couponRepository.FindByCode(input.CouponCode)
		if err != nil {
			return PlaceOrderOutput{}, err
		}
		order.AddCoupon(coupon)
	}
	total := order.GetTotal()
	err = uc.orderRepository.Save(order)
	if err != nil {
		return PlaceOrderOutput{}, err
	}
	return NewPlaceOrderOutput(order.Code.Value, order.ShippingCost, total), nil
}
