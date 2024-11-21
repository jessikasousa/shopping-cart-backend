package repository

import (
	"shopping-cart-backend/internal/model"
	"shopping-cart-backend/pkg/db"
)

func GetCart() ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := db.DB.Find(&cartItems).Error
	return cartItems, err
}

func AddToCart(item model.CartItem) error {
	return db.DB.Create(&item).Error
}

func RemoveFromCart(id uint) error {
	return db.DB.Delete(&model.CartItem{}, id).Error
}