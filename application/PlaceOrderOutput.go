package application

type PlaceOrderOutput struct {
	Code         string
	Total        float64
	ShippingCost float64
}

func NewPlaceOrderOutput(code string, shippingCost float64, total float64) PlaceOrderOutput {
	return PlaceOrderOutput{Code: code, ShippingCost: shippingCost, Total: total}
}
