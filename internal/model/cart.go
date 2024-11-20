package model

type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ProductID string  `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
