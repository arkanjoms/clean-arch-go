package application

import "github.com/google/uuid"

type PlaceOrderItemInput struct {
	ItemID   uuid.UUID
	Quantity float64
}
