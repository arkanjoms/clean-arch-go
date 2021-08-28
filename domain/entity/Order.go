package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrInvalidCPF = errors.New("invalid cpf")

type Order struct {
	ID           uuid.UUID
	Document     Document
	Items        []OrderItem
	Coupon       Coupon
	ShippingCost float64
	Code         OrderCode
	IssueDate    time.Time
	Sequence     int
}

func NewOrder(cpf string, issueDate time.Time, sequence int) (*Order, error) {
	order := Order{
		Document:     NewCPF(cpf),
		Items:        make([]OrderItem, 1),
		ShippingCost: 0,
		Code:         NewOrderCode(issueDate, sequence),
		IssueDate:    issueDate,
		Sequence:     sequence,
	}
	if !order.Document.Valid() {
		return nil, ErrInvalidCPF
	}
	return &order, nil
}

func (o *Order) AddItem(itemID uuid.UUID, price float64, quantity float64) {
	o.Items = append(o.Items, NewOrderItem(itemID, price, quantity))
}

func (o Order) GetTotal() float64 {
	var total float64
	for _, it := range o.Items {
		total += it.Total()
	}
	if (o.Coupon != Coupon{}) {
		total -= (total * o.Coupon.Percentage) / 100
	}
	total += o.ShippingCost
	return total
}

func (o *Order) AddCoupon(coupon Coupon) {
	if coupon.Valid() {
		o.Coupon = coupon
	}
}
