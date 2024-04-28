package repositories

import (
	"context"
	"svc-master/app/models"
)

type shopRepository repository

type ShopRepository interface {
	GetShop(ctx context.Context, filter models.StandardGetRequest) ([]models.Shop, models.Pagination, error)
	GetDetailShop(ctx context.Context, id int) (models.Shop, error)
}

func (r *shopRepository) GetShop(ctx context.Context, filter models.StandardGetRequest) ([]models.Shop, models.Pagination, error) {

	var (
		shops      []models.Shop
		pagination models.Pagination
		totalItems int64
	)

	offset := (filter.Page - 1) * filter.Limit

	query := r.Options.Postgres.Table("shops").Order("shops.name")

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	result := query.Find(&shops)

	// count totalItems by filter
	result = result.Count(&totalItems)
	pagination.Total = int(totalItems)

	result = result.WithContext(ctx).Offset(offset).Limit(filter.Limit).Find(&shops)
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

	return shops, pagination, nil
}

func (r *shopRepository) GetDetailShop(ctx context.Context, id int) (models.Shop, error) {

	var (
		shop models.Shop
	)

	query := r.Options.Postgres.Table("shops").Where("shops.shop_id = ?", id)

	error := query.WithContext(ctx).Find(&shop).Error

	return shop, error
}
