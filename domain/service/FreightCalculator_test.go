package service

import (
	"clean-arch-go/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateItem_value30(t *testing.T) {
	item := entity.Item{ID: uuid.New(), Description: "Item", Price: 1000, Width: 100, Length: 50, Height: 15, Weight: 3}
	distance := 1000.0
	shippingCost := NewFreightCalculator().Calculator(distance, item)
	assert.Equal(t, 30.0, shippingCost)
}

func TestCalculateItem_value220(t *testing.T) {
	item := entity.Item{ID: uuid.New(), Description: "Item", Price: 5000, Width: 50, Length: 50, Height: 50, Weight: 22}
	distance := 1000.0
	shippingCost := NewFreightCalculator().Calculator(distance, item)
	assert.Equal(t, 220.0, shippingCost)
}

func TestCalculateItem_minimumShippingCost10(t *testing.T) {
	item := entity.Item{ID: uuid.New(), Description: "Item", Price: 30, Width: 1, Length: 1, Height: 1, Weight: 1}
	distance := 1000.0
	shippingCost := NewFreightCalculator().Calculator(distance, item)
	assert.Equal(t, 10.0, shippingCost)
}
