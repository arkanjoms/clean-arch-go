package getorder

type OutputGetOrder struct {
	Code         string
	ShippingCost float64
	Total        float64
	OrderItens   []OutputGetOrderItem
}
