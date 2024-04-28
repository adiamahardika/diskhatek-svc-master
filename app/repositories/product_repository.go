package repositories

import (
	"context"
	"svc-master/app/models"
)

type productRepository repository

type ProductRepository interface {
	GetProduct(ctx context.Context, filter models.GetProductRequest) ([]models.Product, models.Pagination, error)
	GetDetailProduct(ctx context.Context, id int) (models.Product, error)
}

func (r *productRepository) GetProduct(ctx context.Context, filter models.GetProductRequest) ([]models.Product, models.Pagination, error) {

	var (
		products   []models.Product
		pagination models.Pagination
		totalItems int64
	)

	offset := (filter.Page - 1) * filter.Limit

	query := r.Options.Postgres.Table("products").Select("products.*, shops.name AS shop, COALESCE(SUM(stocks.quantity),0) - COALESCE(SUM(reserved_stocks.reserved_quantity),0) AS available_stock").Joins("JOIN shops ON products.shop_id = shops.shop_id").Joins("JOIN stocks ON products.product_id = stocks.product_id").Joins("JOIN warehouses ON stocks.warehouse_id = warehouses.warehouse_id AND warehouses.status = 'active'").Joins("LEFT JOIN reserved_stocks ON products.product_id = reserved_stocks.product_id").Group("products.product_id, shops.name").Order("products.name")

	if filter.Name != "" {
		query = query.Where("products.name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.ShopId != 0 {
		query = query.Where("products.shop_id = ?", filter.ShopId)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	result := query.Find(&products)

	// count totalItems by filter
	result = result.Count(&totalItems)
	pagination.Total = int(totalItems)

	result = result.WithContext(ctx).Offset(offset).Limit(filter.Limit).Find(&products)
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

	return products, pagination, nil
}

func (r *productRepository) GetDetailProduct(ctx context.Context, id int) (models.Product, error) {

	var (
		product models.Product
	)

	stockQuery := r.Options.Postgres.Table("stocks").Select("COALESCE(SUM(quantity),0)").Joins("JOIN warehouses ON stocks.warehouse_id = warehouses.warehouse_id AND warehouses.status = 'active'").Where("stocks.product_id = ?", id)
	reservedStockQuery := r.Options.Postgres.Table("reserved_stocks").Select("COALESCE(SUM(reserved_quantity),0)").Where("reserved_stocks.product_id = ?", id)

	query := r.Options.Postgres.Table("products").Select("products.*, shops.name AS shop, (?) - (?) AS available_stock", stockQuery, reservedStockQuery).Joins("JOIN shops ON products.shop_id = shops.shop_id").Where("products.product_id = ?", id).Group("products.product_id, shops.name")

	error := query.WithContext(ctx).Find(&product).Error

	return product, error
}
