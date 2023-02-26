package upoint

import (
	"catalog-product-sa/domain/application"
	"catalog-product-sa/domain/entity"
	"net/http"
)

type upointApplication struct {
	httpClient *http.Client
}

func (u *upointApplication) GetListItemByCodeProduct(codeProduct string) ([]*entity.Item, error) {
	//TODO implement me
	panic("implement me")
}

func NewUpointApplication(client *http.Client) application.ApplicationItem {
	return &upointApplication{httpClient: client}
}
