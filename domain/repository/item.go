package repository

import "catalog-product-sa/domain/entity"

type ItemRepository interface {
	GetItemsByProductSlug(slug string) ([]*entity.Item, error)
}
