package database

import (
	"clean-arch-go/domain/entity"
	"clean-arch-go/domain/repository"
	infraDatabase "clean-arch-go/infra/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ItemRepositoryDatabase struct {
	db infraDatabase.Database
}

func NewItemRepositoryDatabase(db infraDatabase.Database) repository.ItemRepository {
	return ItemRepositoryDatabase{db: db}
}

func (i ItemRepositoryDatabase) GetById(id uuid.UUID) (entity.Item, error) {
	var item entity.Item
	row := i.db.One(context.Background(), "select * from ccca.item where id = $1", id)
	err := row.Scan(&item.ID, &item.Category, &item.Description, &item.Price, &item.Width, &item.Height, &item.Length, &item.Weight)
	if err != nil && !errors.Is(err, infraDatabase.ErrNoRows) {
		return entity.Item{}, fmt.Errorf("could not scan result into item: %w", err)
	}
	return item, nil
}

func (i ItemRepositoryDatabase) FindAll() ([]entity.Item, error) {
	rows, err := i.db.Many(context.Background(), "select id, category, description, price from ccca.item")
	if err != nil && !errors.Is(err, infraDatabase.ErrNoRows) {
		return nil, fmt.Errorf("could not find item's: %w", err)
	}
	if errors.Is(err, infraDatabase.ErrNoRows) {
		return nil, fmt.Errorf("no items found")
	}
	return i.parseResult(rows)
}

func (i ItemRepositoryDatabase) FindAllByCategory(category string) ([]entity.Item, error) {
	rows, err := i.db.Many(context.Background(), "select id, category, description, price from ccca.item where category = $1", category)
	if err != nil && !errors.Is(err, infraDatabase.ErrNoRows) {
		return nil, fmt.Errorf("could not find item's: %w", err)
	}
	if errors.Is(err, infraDatabase.ErrNoRows) {
		return nil, fmt.Errorf("no items found")
	}
	return i.parseResult(rows)
}

func (i ItemRepositoryDatabase) parseResult(rows *sql.Rows) (items []entity.Item, err error) {
	if errors.Is(err, infraDatabase.ErrNoRows) {
		return nil, fmt.Errorf("no items found")
	}
	for rows.Next() {
		var item entity.Item
		err = rows.Scan(&item.ID, &item.Category, &item.Description, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("could not scan item: %w", err)
		}
		items = append(items, item)
	}
	return items, nil
}
