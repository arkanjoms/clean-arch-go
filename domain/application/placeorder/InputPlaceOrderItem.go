package placeorder

import "github.com/google/uuid"

type InputPlaceOrderItem struct {
	ItemID   uuid.UUID
	Quantity float64
}
