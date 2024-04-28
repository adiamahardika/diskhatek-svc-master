package repositories

import (
	"svc-master/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	Shop      ShopRepository
	Warehouse WarehouseRepository
}

type repository struct {
	Options Options
}

type Options struct {
	Postgres *gorm.DB
	Config   *config.Config
}

func Init(opts Options) *Main {
	repo := &repository{opts}

	m := &Main{
		Shop:      (*shopRepository)(repo),
		Warehouse: (*warehouseRepository)(repo),
	}

	return m
}
