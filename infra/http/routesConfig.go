package http

import (
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/gateway"
	"clean-arch-go/domain/service"
	"clean-arch-go/infra/http/handler"
)

type RoutesConfig struct {
	server            Http
	repositoryFactory factory.RepositoryFactory
	zipcodeClient     gateway.ZipcodeClient
	freightCalculator service.FreightCalculator
}

func NewRouteConfig(server Http, repositoryFactory factory.RepositoryFactory, zipcodeClient gateway.ZipcodeClient, freightCalculator service.FreightCalculator) RoutesConfig {
	return RoutesConfig{
		server:            server,
		repositoryFactory: repositoryFactory,
		zipcodeClient:     zipcodeClient,
		freightCalculator: freightCalculator,
	}
}

func (c RoutesConfig) Build() {
	c.server.On("get", "/orders/${code}", handler.NewOrderHandler(c.repositoryFactory).GetOrder)
	c.server.On("get", "/items", handler.NewItemHandler(c.repositoryFactory).GetItems)
	c.server.On("post", "/orders", handler.NewPlaceOrderHandler(c.zipcodeClient, c.freightCalculator, c.repositoryFactory).CreateOrder)
}
