package usecases

import (
	"context"
	"svc-master/app/models"
)

type productUsecase usecase

type ProductUsecase interface {
	GetProduct(ctx context.Context, filter models.GetProductRequest) ([]models.Product, models.Pagination, error)
}

func (u *productUsecase) GetProduct(ctx context.Context, filter models.GetProductRequest) ([]models.Product, models.Pagination, error) {
	var (
		products   []models.Product
		pagination models.Pagination
		err        error
	)

	if filter.Limit == 0 {
		filter.Limit = 10
	}

	if filter.Page == 0 {
		filter.Page = 1
	}

	products, pagination, err = u.Options.Repository.Product.GetProduct(ctx, filter)

	return products, pagination, err
}
