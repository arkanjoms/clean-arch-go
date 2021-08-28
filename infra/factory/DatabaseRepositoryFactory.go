package factory

import (
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/repository"
	pgDatabase "clean-arch-go/infra/database"
	"clean-arch-go/infra/repository/database"
)

type DatabaseRepositoryFactory struct {
	db *pgDatabase.PGDatabase
}

func NewDatabaseRepositoryFactory(db *pgDatabase.PGDatabase) factory.RepositoryFactory {
	return DatabaseRepositoryFactory{db: db}
}

func (f DatabaseRepositoryFactory) NewItemRepository() repository.ItemRepository {
	return database.NewItemRepositoryDatabase(f.db.GetInstance())
}

func (f DatabaseRepositoryFactory) NewCouponRepository() repository.CouponRepository {
	return database.NewCouponRepositoryDatabase(f.db.GetInstance())
}

func (f DatabaseRepositoryFactory) NewOrderRepository() repository.OrderRepository {
	return database.NewOrderRepositoryDatabase(f.db.GetInstance())
}
