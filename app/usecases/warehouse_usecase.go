package usecases

import (
	"context"
	"svc-master/app/models"
	"time"
)

type warehouseUsecase usecase

type WarehouseUsecase interface {
	GetWarehouse(ctx context.Context, filter models.GetWarehouseRequest) ([]models.Warehouse, models.Pagination, error)
	UpdateWarehouse(request models.Warehouse, id int) (models.Warehouse, error)
}

func (u *warehouseUsecase) GetWarehouse(ctx context.Context, filter models.GetWarehouseRequest) ([]models.Warehouse, models.Pagination, error) {
	var (
		warehouses []models.Warehouse
		pagination models.Pagination
		err        error
	)

	if filter.Limit == 0 {
		filter.Limit = 10
	}

	if filter.Page == 0 {
		filter.Page = 1
	}

	warehouses, pagination, err = u.Options.Repository.Warehouse.GetWarehouse(ctx, filter)

	return warehouses, pagination, err
}

func (u *warehouseUsecase) UpdateWarehouse(request models.Warehouse, id int) (models.Warehouse, error) {
	var (
		err       error
		warehouse models.Warehouse
	)

	request.UpdatedAt = time.Now()

	warehouse, err = u.Options.Repository.Warehouse.UpdateWarehouse(request, id)
	return warehouse, err
}
