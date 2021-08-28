package memory

import (
	"clean-arch-go/domain/gateway"
	"fmt"
)

type ZipcodeClientMemory struct{}

func NewZipcodeClient() gateway.ZipcodeClient {
	return ZipcodeClientMemory{}
}

func (z ZipcodeClientMemory) Distance(origin string, destiny string) float64 {
	fmt.Printf("Distance from %s to %s\n", origin, destiny)
	return 1000
}
