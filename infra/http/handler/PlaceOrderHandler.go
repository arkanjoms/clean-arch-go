package handler

import (
	"clean-arch-go/application/placeorder"
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/gateway"
	"clean-arch-go/domain/service"
	"encoding/json"
	"io"
)

type PlaceOrderHandler struct {
	zipcodeClient     gateway.ZipcodeClient
	freightCalculator service.FreightCalculator
	repositoryFactory factory.RepositoryFactory
}

func NewPlaceOrderHandler(zipcodeClient gateway.ZipcodeClient, freightCalculator service.FreightCalculator, repositoryFactory factory.RepositoryFactory) PlaceOrderHandler {
	return PlaceOrderHandler{
		zipcodeClient:     zipcodeClient,
		freightCalculator: freightCalculator,
		repositoryFactory: repositoryFactory,
	}
}

func (h PlaceOrderHandler) CreateOrder(_ map[string]string, _ map[string][]string, body io.ReadCloser) (interface{}, error) {
	placeOrder := placeorder.NewPlaceOrder(h.zipcodeClient, h.freightCalculator, h.repositoryFactory)
	var input placeorder.InputPlaceOrder
	_ = json.NewDecoder(body).Decode(&input)
	return placeOrder.Execute(input)
}
