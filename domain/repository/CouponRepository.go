package repository

import (
	"clean-arch-go/domain/entity"
)

type CouponRepository interface {
	FindByCode(code string) (entity.Coupon, error)
}
