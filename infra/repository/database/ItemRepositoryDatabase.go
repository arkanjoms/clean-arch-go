package database

import (
	"clean-arch-go/domain/entity"
	"clean-arch-go/domain/repository"
	infraDB "clean-arch-go/infra/database"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type ItemRepositoryDatabase struct {
	db infraDB.Database
}

func NewItemRepositoryDatabase(db infraDB.Database) repository.ItemRepository {
	return ItemRepositoryDatabase{db: db}
}

func (i ItemRepositoryDatabase) GetById(id uuid.UUID) (entity.Item, error) {
	var item entity.Item
	row := i.db.One(context.Background(), "select * from ccca.item where id = $1", id)
	err := row.Scan(&item.ID, &item.Category, &item.Description, &item.Price, &item.Width, &item.Height, &item.Length, &item.Weight)
	if err != nil {
		return entity.Item{}, fmt.Errorf("could not scan result into item: %w", err)
	}
	return item, nil
}
