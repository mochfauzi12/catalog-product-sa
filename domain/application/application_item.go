package application

import "catalog-product-sa/domain/entity"

type ApplicationItem interface {
	GetListItemByCodeProduct(codeProduct string) ([]*entity.Item, error)
}
