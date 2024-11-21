package service

import (
	"shopping-cart-backend/internal/model"
	"shopping-cart-backend/internal/repository"

	"shopping-cart-backend/pkg/db"
)

func GetCart() ([]model.CartItem, error) {
	var cart []model.CartItem
	err := db.DB.Find(&cart).Error 
	return cart, err
}

func AddToCart(item model.CartItem) error {
	return repository.AddToCart(item)
}

func RemoveFromCart(id uint) error {
	return db.DB.Delete(&model.CartItem{}, id).Error
}

func RemoveItemFromCart(id uint) error {
	return repository.RemoveFromCart(id)
}

func AddItemToCart(productID, title string, quantity int, price float64) error {
	cartItem := model.CartItem{
		ProductID: productID,
		Title:     title,
		Quantity:  quantity,
		Price:     price,
	}
	return db.DB.Create(&cartItem).Error
}