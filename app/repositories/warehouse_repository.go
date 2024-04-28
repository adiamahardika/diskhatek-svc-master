package repositories

import (
	"context"
	"svc-master/app/models"
)

type warehouseRepository repository

type WarehouseRepository interface {
	GetWarehouse(ctx context.Context, filter models.GetWarehouseRequest) ([]models.Warehouse, models.Pagination, error)
}

func (r *warehouseRepository) GetWarehouse(ctx context.Context, filter models.GetWarehouseRequest) ([]models.Warehouse, models.Pagination, error) {

	var (
		warehouses []models.Warehouse
		pagination models.Pagination
		totalItems int64
	)

	offset := (filter.Page - 1) * filter.Limit

	query := r.Options.Postgres.Table("warehouses").Select("warehouses.*, shops.name AS shop").Joins("JOIN shops ON warehouses.shop_id = shops.shop_id").Order("warehouses.name")

	if filter.Name != "" {
		query = query.Where("warehouses.name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.ShopId != 0 {
		query = query.Where("warehouses.shop_id = ?", filter.ShopId)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	result := query.Find(&warehouses)

	// count totalItems by filter
	result = result.Count(&totalItems)
	pagination.Total = int(totalItems)

	result = result.WithContext(ctx).Offset(offset).Limit(filter.Limit).Find(&warehouses)
	if result.Error != nil {
		return nil, pagination, result.Error
	}

	// Calculate the total number of pages
	if totalItems%int64(filter.Limit) == 0 {
		pagination.TotalPage = int(totalItems / int64(filter.Limit))
	} else {
		pagination.TotalPage = int(totalItems/int64(filter.Limit)) + 1
	}

	pagination.Page = filter.Page
	pagination.PageSize = filter.Limit

	return warehouses, pagination, nil
}
