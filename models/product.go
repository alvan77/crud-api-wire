package models

import "time"

type Product struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code" validate:"required"`
	Price     uint      `json:"price" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
