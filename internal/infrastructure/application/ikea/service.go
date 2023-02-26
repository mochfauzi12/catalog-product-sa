package ikea

import (
	"catalog-product-sa/domain/application"
	"catalog-product-sa/domain/entity"
	"net/http"
)

type ikeaApplicationInteractor struct {
	httpClient *http.Client
}

func (i *ikeaApplicationInteractor) GetListItemByCodeProduct(codeProduct string) ([]*entity.Item, error) {
	//TODO implement me
	panic("implement me")
}

func NewIkeaApplication(client *http.Client) application.ApplicationItem {
	return &ikeaApplicationInteractor{httpClient: client}
}
