package models

import "time"

type GetProductRequest struct {
	StandardGetRequest
	ShopId int `json:"shop_id"`
}

type Product struct {
	ProductId   int       `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	TotalStocks int       `json:"total_stocks"`
	ShopId      int       `json:"shop_id"`
	Shop        string    `json:"shop"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
