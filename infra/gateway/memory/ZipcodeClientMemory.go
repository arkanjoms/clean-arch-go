package memory

import (
	"clean-arch-go/domain/gateway"
	"github.com/sirupsen/logrus"
)

type ZipcodeClientMemory struct{}

func NewZipcodeClient() gateway.ZipcodeClient {
	return ZipcodeClientMemory{}
}

func (z ZipcodeClientMemory) Distance(origin string, destiny string) float64 {
	logrus.Warnf("Distance from %s to %s", origin, destiny)
	return 1000
}
