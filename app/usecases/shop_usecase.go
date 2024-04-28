package usecases

import (
	"context"
	"svc-master/app/models"
)

type shopUsecase usecase

type ShopUsecase interface {
	GetShop(ctx context.Context, filter models.StandardGetRequest) ([]models.Shop, models.Pagination, error)
}

func (u *shopUsecase) GetShop(ctx context.Context, filter models.StandardGetRequest) ([]models.Shop, models.Pagination, error) {
	var (
		shops      []models.Shop
		pagination models.Pagination
		err        error
	)

	if filter.Limit == 0 {
		filter.Limit = 10
	}

	if filter.Page == 0 {
		filter.Page = 1
	}

	shops, pagination, err = u.Options.Repository.Shop.GetShop(ctx, filter)

	return shops, pagination, err
}
