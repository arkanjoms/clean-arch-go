package placeorder

import "time"

type InputPlaceOrder struct {
	Document       string
	Items          []InputPlaceOrderItem
	CouponCode     string
	ZipcodeOrigin  string
	ZipcodeDestiny string
	IssueDate      time.Time
}
