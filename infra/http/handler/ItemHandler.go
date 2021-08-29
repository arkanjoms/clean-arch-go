package handler

import (
	"clean-arch-go/application"
	"clean-arch-go/domain/factory"
	"io"
)

type ItemHandler struct {
	repositoryFactory factory.RepositoryFactory
}

func NewItemHandler(repositoryFactory factory.RepositoryFactory) ItemHandler {
	return ItemHandler{repositoryFactory: repositoryFactory}
}

func (h ItemHandler) GetItems(_ map[string]string, queryParams map[string][]string, _ io.ReadCloser) (interface{}, error) {
	getItem := application.NewGetItem(h.repositoryFactory)
	category := getQueryParamOne(queryParams, "category")
	return getItem.Execute(category)
}
