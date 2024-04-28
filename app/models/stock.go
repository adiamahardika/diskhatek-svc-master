package models

import "time"

type Stock struct {
	StockId     int       `json:"stock_id" gorm:"primaryKey"`
	ProductId   int       `json:"product_id"`
	WarehouseId int       `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
