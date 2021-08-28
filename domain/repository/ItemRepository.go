package repository

import (
	"clean-arch-go/domain/entity"
	"github.com/google/uuid"
)

type ItemRepository interface {
	GetById(id uuid.UUID) (entity.Item, error)
}
