package getitem

import "github.com/google/uuid"

type OutputGetItem struct {
	ID          uuid.UUID
	Category    string
	Description string
	Price       float64
}
