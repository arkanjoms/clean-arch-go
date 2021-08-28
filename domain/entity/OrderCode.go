package entity

import (
	"fmt"
	"time"
)

type OrderCode struct {
	Value string
}

func NewOrderCode(issueDate time.Time, sequence int) OrderCode {
	return OrderCode{Value: fmt.Sprintf("%d%08d", issueDate.Year(), sequence)}
}
