package product

import (
	"catalog-product-sa/domain/repository"
	"catalog-product-sa/domain/service"
	"catalog-product-sa/domain/usecase"
)

type productUseCaseInteractor struct {
	productService service.ProductService
	productRepo    repository.ProductRepository
}

func NewProductUseCase(concreteProductService service.ProductService,
	repoProduct repository.ProductRepository) usecase.ProductUseCase {
	return &productUseCaseInteractor{
		productService: concreteProductService,
		productRepo:    repoProduct,
	}
}
