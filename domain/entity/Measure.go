package entity

type Measure float64

func (m Measure) InMeters() float64 {
	return float64(m) / 100
}
