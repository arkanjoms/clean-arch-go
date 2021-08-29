package repository

import (
	"clean-arch-go/domain/entity"
	"github.com/google/uuid"
)

type ItemRepository interface {
	GetById(id uuid.UUID) (entity.Item, error)
	FindAll() ([]entity.Item, error)
	FindAllByCategory(category string) ([]entity.Item, error)
}
