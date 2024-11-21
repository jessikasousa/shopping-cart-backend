package model

import "time"

type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID string    `gorm:"not null" json:"product_id"`
	Title     string    `gorm:"not null" json:"title"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
