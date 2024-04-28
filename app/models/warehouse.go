package models

import "time"

type GetWarehouseRequest struct {
	StandardGetRequest
	ShopId int `json:"shop_id"`
}

type Warehouse struct {
	WarehouseId int       `json:"warehouse_id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	ShopId      int       `json:"shop_id"`
	Shop        string    `json:"shop"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
