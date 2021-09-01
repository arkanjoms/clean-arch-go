package factory

import (
	"clean-arch-go/domain/factory"
	"clean-arch-go/domain/repository"
	"clean-arch-go/infra/database"
	repoDatabase "clean-arch-go/infra/repository/database"
)

type DatabaseRepositoryFactory struct {
	pgDB *database.PGDatabase
}

func NewDatabaseRepositoryFactory(pgDB *database.PGDatabase) factory.RepositoryFactory {
	return DatabaseRepositoryFactory{pgDB: pgDB}
}

func (f DatabaseRepositoryFactory) NewItemRepository() repository.ItemRepository {
	return repoDatabase.NewItemRepositoryDatabase(f.pgDB)
}

func (f DatabaseRepositoryFactory) NewCouponRepository() repository.CouponRepository {
	return repoDatabase.NewCouponRepositoryDatabase(f.pgDB)
}

func (f DatabaseRepositoryFactory) NewOrderRepository() repository.OrderRepository {
	return repoDatabase.NewOrderRepositoryDatabase(f.pgDB)
}
