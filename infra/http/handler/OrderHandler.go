package handler

import (
	"clean-arch-go/application"
	"clean-arch-go/domain/factory"
	"io"
)

type OrderHandler struct {
	repositoryFactory factory.RepositoryFactory
}

func NewOrderHandler(repositoryFactory factory.RepositoryFactory) OrderHandler {
	return OrderHandler{repositoryFactory: repositoryFactory}
}

func (h OrderHandler) GetOrder(params map[string]string, _ map[string][]string, _ io.ReadCloser) (interface{}, error) {
	getOrder := application.NewGetOrder(h.repositoryFactory)
	return getOrder.Execute(params["code"])
}
