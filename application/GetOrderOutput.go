package application

type GetOrderOutput struct {
	Code         string
	ShippingCost float64
	Total        float64
	OrderItens   []GetOrderItemOutput
}
