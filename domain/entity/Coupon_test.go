package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCouponIsExpired(t *testing.T) {
	coupon := NewCoupon("CODE10", 10.00, time.Now().Add(-1*time.Hour))
	assert.False(t, coupon.Valid())
}

func TestCouponIsValid(t *testing.T) {
	coupon := NewCoupon("CODE10", 10.00, time.Now().Add(1*time.Hour))
	assert.True(t, coupon.Valid())
}
