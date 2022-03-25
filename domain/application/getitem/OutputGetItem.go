package getitem

import "github.com/google/uuid"

type OutputGetItem struct {
	ID          uuid.UUID `json:"id"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}
