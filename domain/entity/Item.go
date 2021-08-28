package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID
	Category    string
	Description string
	Price       float64
	Width       Measure
	Length      Measure
	Height      Measure
	Weight      float64
}

func NewItem(id uuid.UUID, description string, price float64, width Measure, length Measure, height Measure, weight float64) Item {
	return Item{ID: id, Description: description, Price: price, Width: width, Length: length, Height: height, Weight: weight}
}

func (i Item) Volume() float64 {
	return i.Height.InMeters() * i.Width.InMeters() * i.Length.InMeters()
}

func (i Item) Density() float64 {
	return i.Weight / i.Volume()
}
