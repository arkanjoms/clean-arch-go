package application

import "github.com/google/uuid"

type GetItemOutput struct {
	ID          uuid.UUID
	Category    string
	Description string
	Price       float64
}
