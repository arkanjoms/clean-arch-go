package gateway

type ZipcodeClient interface {
	Distance(origin string, destiny string) float64
}
