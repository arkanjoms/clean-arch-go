package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDoNotCreateOrderWithInvalidCPF(t *testing.T) {
	cpf := "111.111.111-11"
	order, err := NewOrder(cpf, time.Now(), 1)
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.ErrorAs(t, ErrInvalidCPF, &err)
}

func TestCreateOrderWith3Itens(t *testing.T) {
	cpf := "115.781.529-41"
	order, err := NewOrder(cpf, time.Now(), 1)
	assert.NoError(t, err)
	order.AddItem(uuid.New(), 1000, 2)
	order.AddItem(uuid.New(), 5000, 1)
	order.AddItem(uuid.New(), 30, 3)
	total := order.GetTotal()
	assert.Equal(t, float64(7090), total)
}

func TestCreateOrderWith3ItensWithCoupon(t *testing.T) {
	cpf := "115.781.529-41"
	order, err := NewOrder(cpf, time.Now(), 1)
	assert.NoError(t, err)
	order.AddItem(uuid.New(), 1000, 2)
	order.AddItem(uuid.New(), 5000, 1)
	order.AddItem(uuid.New(), 30, 3)
	order.AddCoupon(NewCoupon("CODE20", 20.0, time.Now().Add(1*time.Hour)))
	total := order.GetTotal()
	assert.Equal(t, float64(5672), total)
}
