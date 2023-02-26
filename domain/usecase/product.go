package usecase

import "catalog-product-sa/domain/entity"

type ProductUseCase interface {
	GetDetailProduct(slug string) (*entity.Product, error)
}
