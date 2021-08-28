package application

import (
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/repository"
)

type GetOrder struct {
	itemRepository   repository.ItemRepository
	couponRepository repository.CouponRepository
	orderRepository  repository.OrderRepository
}

func NewGetOrder(repositoryFactory factory.RepositoryFactory) *GetOrder {
	return &GetOrder{
		itemRepository:   repositoryFactory.NewItemRepository(),
		couponRepository: repositoryFactory.NewCouponRepository(),
		orderRepository:  repositoryFactory.NewOrderRepository(),
	}
}

func (s GetOrder) Execute(code string) (*GetOrderOutput, error) {
	order, err := s.orderRepository.Get(code)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, nil
	}
	outputOrder := &GetOrderOutput{
		Code:         order.Code.Value,
		ShippingCost: order.ShippingCost,
		Total:        order.GetTotal(),
		OrderItens:   make([]GetOrderItemOutput, 0, len(order.Items)),
	}
	for _, orderItem := range order.Items {
		item, err := s.itemRepository.GetById(orderItem.ItemID)
		if err != nil {
			return nil, err
		}
		outputOrder.OrderItens = append(outputOrder.OrderItens, GetOrderItemOutput{
			Description: item.Description,
			Price:       orderItem.Price,
			Quantity:    orderItem.Quantity,
		})
	}
	return outputOrder, nil
}
