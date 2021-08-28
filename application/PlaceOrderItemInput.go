package application

import "github.com/google/uuid"

type PlaceOrderItemInput struct {
	itemID   uuid.UUID
	quantity float64
}
