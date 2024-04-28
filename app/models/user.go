package models

import "time"

type GetUserRequest struct {
	StandardGetRequest
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type User struct {
	UserId       int       `json:"user_id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Password     string    `json:"password,omitempty" gorm:"-:all"`
	PasswordHash string    `json:"password_hash,omitempty" gorm:"->:false;<-:create"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
