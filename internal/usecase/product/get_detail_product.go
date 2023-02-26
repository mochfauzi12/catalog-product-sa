package product

import (
	"catalog-product-sa/domain/entity"
	"errors"
)

func (pro *productUseCaseInteractor) GetDetailProduct(slug string) (*entity.Product, error) {
	product, err := pro.productRepo.GetProductBySlug(slug)
	if err != nil {
		return nil, errors.New("PRODUCT NOT FOUND!")
	}

	productWithItems, errGetItems := pro.productService.GetListItemOfProduct(product)
	if errGetItems != nil {
		return nil, errors.New("ITEM NOT FOUND!")
	}

	return productWithItems, nil
}
