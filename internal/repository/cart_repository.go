package repository

import (
	"shopping-cart-backend/internal/model"
	"shopping-cart-backend/pkg/db"
)

func GetCart() ([]model.CartItem, error) {
	var cart []model.CartItem
	err := db.DB.Find(&cart).Error
	return cart, err
}

func AddToCart(item model.CartItem) error {
	return db.DB.Create(&item).Error
}

func RemoveFromCart(id uint) error {
	return db.DB.Delete(&model.CartItem{}, id).Error
}
