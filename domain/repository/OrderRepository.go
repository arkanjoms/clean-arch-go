package repository

import (
	"clean-arch-go/domain/entity"
)

type OrderRepository interface {
	Save(order *entity.Order) error
	Get(code string) (*entity.Order, error)
	Count() int
}
