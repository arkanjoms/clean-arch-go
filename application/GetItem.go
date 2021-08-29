package application

import (
	"clean-arch-go/domain/entity"
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/repository"
)

type GetItem struct {
	itemRepository repository.ItemRepository
}

func NewGetItem(repositoryFactory factory.RepositoryFactory) GetItem {
	return GetItem{itemRepository: repositoryFactory.NewItemRepository()}
}

func (i GetItem) Execute(category string) ([]GetItemOutput, error) {
	items, err := i.findItems(category)
	if err != nil {
		return nil, err
	}
	var itemsOutput []GetItemOutput
	for _, item := range items {
		itemsOutput = append(itemsOutput, GetItemOutput{
			ID:          item.ID,
			Category:    item.Category,
			Description: item.Description,
			Price:       item.Price,
		})
	}
	return itemsOutput, nil
}

func (i GetItem) findItems(category string) ([]entity.Item, error) {
	if category == "" {
		return i.itemRepository.FindAll()
	}
	return i.itemRepository.FindAllByCategory(category)
}
