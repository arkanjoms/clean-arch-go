package entity

const (
	factorDigit1 = 10
	factorDigit2 = 11
	maxDigits1   = 9
	maxDigits2   = 10
)

type Document interface {
	Valid() bool
	Value() string
}
