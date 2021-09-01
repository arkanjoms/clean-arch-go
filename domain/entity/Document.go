package entity

type Document interface {
	Valid() bool
	Value() string
}
