package usecases

import (
	"svc-master/app/repositories"
	"svc-master/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	Shop          ShopUsecase
	Warehouse     WarehouseUsecase
	Product       ProductUsecase
	StockTransfer StockTransferUsecase
	Validate      ValidateUsecase
}

type usecase struct {
	Options Options
}

type Options struct {
	Repository *repositories.Main
	Config     *config.Config
	Postgres   *gorm.DB
}

func Init(opts Options) *Main {
	uscs := &usecase{opts}

	m := &Main{
		Shop:          (*shopUsecase)(uscs),
		Warehouse:     (*warehouseUsecase)(uscs),
		Product:       (*productUsecase)(uscs),
		StockTransfer: (*stockTransferUsecase)(uscs),
		Validate:      (*validateUsecase)(uscs),
	}

	return m
}
