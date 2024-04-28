package models

import "time"

type StockTransfer struct {
	TransferId      int       `json:"transfer_id" gorm:"primaryKey"`
	ProductId       int       `json:"product_id"`
	SrcWarehouseId  int       `json:"source_warehouse_id" gorm:"column:source_warehouse_id"`
	DestWarehouseId int       `json:"destination_warehouse_id" gorm:"column:destination_warehouse_id"`
	Quantity        int       `json:"quantity"`
	TransferDate    string    `json:"transfer_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
