//+build wireinject

package main

import (
	"crud-api-wire/api"
	"crud-api-wire/repository"
	"crud-api-wire/service"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func ProductHandler(db *gorm.DB) api.ProductAPI {
	wire.Build(repository.ProvideProductRepository, service.ProvideProductService, api.ProvideProductAPI)
	return api.ProductAPI{}
}
