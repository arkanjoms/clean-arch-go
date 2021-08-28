package application

import "time"

type PlaceOrderInput struct {
	Document       string
	Items          []PlaceOrderItemInput
	CouponCode     string
	ZipcodeOrigin  string
	ZipcodeDestiny string
	IssueDate      time.Time
}
