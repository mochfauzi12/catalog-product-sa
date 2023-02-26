package product

import (
	"catalog-product-sa/domain/application"
	"catalog-product-sa/domain/repository"
	"catalog-product-sa/domain/service"
)

type productServiceInteractor struct {
	itemRepository repository.ItemRepository
	appRepository  application.ApplicationItem
}

func NewProductServiceInteractor(concreteItemRepo repository.ItemRepository,
	concreteApplicationRepo application.ApplicationItem) service.ProductService {
	return &productServiceInteractor{
		itemRepository: concreteItemRepo,
		appRepository:  concreteApplicationRepo,
	}
}
