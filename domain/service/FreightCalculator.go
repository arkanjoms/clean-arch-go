package service

import (
	"clean-arch-go/domain/entity"
	"math"
)

const MinimumShippingCost = 10.0

type FreightCalculator struct {
}

func NewFreightCalculator() FreightCalculator {
	return FreightCalculator{}
}

func (f FreightCalculator) Calculator(distance float64, item entity.Item) float64 {
	shippingCost := distance * item.Volume() * (math.Floor(item.Density()) / 100.00)
	if shippingCost < MinimumShippingCost {
		return MinimumShippingCost
	}
	return shippingCost
}
