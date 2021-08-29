package database

import (
	"clean-arch-go/domain/entity"
	"clean-arch-go/domain/repository"
	infraDatabase "clean-arch-go/infra/database"
	"context"
	"errors"
	"fmt"
)

type CouponRepositoryDatabase struct {
	db infraDatabase.Database
}

func NewCouponRepositoryDatabase(db infraDatabase.Database) repository.CouponRepository {
	return CouponRepositoryDatabase{db: db}
}

func (c CouponRepositoryDatabase) FindByCode(code string) (entity.Coupon, error) {
	coupon := &entity.Coupon{}
	row := c.db.One(context.Background(), "select code, percentage, expire_date from ccca.coupon where code = $1", code)
	err := row.Scan(&coupon.Code, &coupon.Percentage, &coupon.ExpiresIn)
	if err != nil && !errors.Is(err, infraDatabase.ErrNoRows) {
		return entity.Coupon{}, fmt.Errorf("could not scan coupon: %w", err)
	}
	return *coupon, nil
}
