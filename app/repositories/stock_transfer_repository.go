package repositories

import (
	"svc-master/app/models"

	"gorm.io/gorm"
)

type stockTransferRepository repository

type StockTransferRepository interface {
	CreateStockTransfer(tx *gorm.DB, request models.StockTransfer) (models.StockTransfer, error)
}

func (r *stockTransferRepository) CreateStockTransfer(tx *gorm.DB, request models.StockTransfer) (models.StockTransfer, error) {

	err := tx.Save(&request).Error

	return request, err
}
