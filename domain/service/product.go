package service

import "catalog-product-sa/domain/entity"

type ProductService interface {
	GetListItemOfProduct(product *entity.Product) (*entity.Product, error)
}
