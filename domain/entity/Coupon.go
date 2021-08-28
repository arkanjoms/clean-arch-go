package entity

import (
	"time"
)

type Coupon struct {
	Code       string
	Percentage float64
	ExpiresIn  time.Time
}

func NewCoupon(code string, percentage float64, expiresIn time.Time) Coupon {
	return Coupon{Code: code, Percentage: percentage, ExpiresIn: expiresIn}
}

func (c Coupon) Valid() bool {
	return time.Now().Before(c.ExpiresIn)
}
