package usecases

import (
	"svc-master/app/models"
	"time"
)

type stockTransferUsecase usecase

type StockTransferUsecase interface {
	CreateStockTransfer(request models.StockTransfer) (models.StockTransfer, error)
}

func (u *stockTransferUsecase) CreateStockTransfer(request models.StockTransfer) (models.StockTransfer, error) {

	now := time.Now()

	tx := u.Options.Postgres.Begin()
	result, err := u.Options.Repository.StockTransfer.CreateStockTransfer(tx, request)
	if err != nil {
		tx.Rollback()
		return models.StockTransfer{}, err
	}

	err = u.Options.Repository.Stock.ReduceStock(tx, models.Stock{
		ProductId:   request.ProductId,
		WarehouseId: request.SrcWarehouseId,
		Quantity:    request.Quantity,
		UpdatedAt:   now,
	})
	if err != nil {
		tx.Rollback()
		return models.StockTransfer{}, err
	}

	err = u.Options.Repository.Stock.AddStock(tx, models.Stock{
		ProductId:   request.ProductId,
		WarehouseId: request.DestWarehouseId,
		Quantity:    request.Quantity,
		UpdatedAt:   now,
	})
	if err != nil {
		tx.Rollback()
		return models.StockTransfer{}, err
	}

	tx.Commit()
	return result, nil
}
