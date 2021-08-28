package entity

import "github.com/google/uuid"

type OrderItem struct {
	ItemID   uuid.UUID
	Price    float64
	Quantity float64
}

func NewOrderItem(itemId uuid.UUID, price float64, quantity float64) OrderItem {
	return OrderItem{
		ItemID:   itemId,
		Price:    price,
		Quantity: quantity,
	}
}

func (i OrderItem) Total() float64 {
	return i.Price * i.Quantity
}
