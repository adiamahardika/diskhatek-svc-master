package models

import "time"

type ReservedStock struct {
	ReservationId         int       `json:"reservation_id" gorm:"primaryKey"`
	OrderItemId           int       `json:"order_item_id"`
	ProductId             int       `json:"product_id"`
	ReservedQuantity      int       `json:"reserverd_quantity"`
	ReservationExpiryTime time.Time `json:"reservation_expiry_time"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
