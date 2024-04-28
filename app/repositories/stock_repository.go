package repositories

import (
	"svc-master/app/models"

	"gorm.io/gorm"
)

type stockRepository repository

type StockRepository interface {
	AddStock(tx *gorm.DB, request models.Stock) error
	ReduceStock(tx *gorm.DB, request models.Stock) error
}

func (r *stockRepository) AddStock(tx *gorm.DB, request models.Stock) error {

	stock := models.Stock{}
	err := tx.Table("stocks").Where("product_id = ? AND warehouse_id = ?", request.ProductId, request.WarehouseId).Find(&stock).Error
	if err != nil {
		return err
	}

	if stock.WarehouseId == 0 {
		stock.WarehouseId = request.WarehouseId
	}
	if stock.ProductId == 0 {
		stock.ProductId = request.ProductId
	}
	stock.Quantity = stock.Quantity + request.Quantity
	stock.UpdatedAt = request.UpdatedAt

	err = tx.Save(&stock).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *stockRepository) ReduceStock(tx *gorm.DB, request models.Stock) error {

	stock := models.Stock{}
	err := tx.Table("stocks").Where("product_id = ? AND warehouse_id = ?", request.ProductId, request.WarehouseId).Find(&stock).Error
	if err != nil {
		return err
	}

	stock.Quantity = stock.Quantity - request.Quantity
	stock.UpdatedAt = request.UpdatedAt

	err = tx.Save(&stock).Error
	if err != nil {
		return err
	}

	return nil
}
