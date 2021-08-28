package entity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestItemGetVolume(t *testing.T) {
	item := NewItem(uuid.New(), "Camera", 1000, 15, 20, 10, 1)
	assert.Equal(t, 0.003, item.Volume())
}

func TestItemDensity(t *testing.T) {
	item := NewItem(uuid.New(), "Camera", 1000, 15, 20, 10, 1)
	density, err := strconv.ParseFloat(fmt.Sprintf("%.2f", item.Density()), 64)
	assert.NoError(t, err)
	assert.Equal(t, 333.33, density)
}
