package factory

import "clean-arch-go/domain/repository"

type RepositoryFactory interface {
	NewItemRepository() repository.ItemRepository
	NewCouponRepository() repository.CouponRepository
	NewOrderRepository() repository.OrderRepository
}
