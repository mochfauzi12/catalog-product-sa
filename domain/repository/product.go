package repository

import "catalog-product-sa/domain/entity"

type ProductRepository interface {
	GetAllProduct() ([]*entity.Product, error)
	GetProductBySlug(slug string) (*entity.Product, error)
}
