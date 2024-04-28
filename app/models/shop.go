package models

import "time"

type Shop struct {
	ShopId    int       `json:"shop_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
