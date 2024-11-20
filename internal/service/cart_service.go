package service

import (
	"shopping-cart-backend/internal/model"
	"shopping-cart-backend/internal/repository"
)

func GetCart() ([]model.CartItem, error) {
	return repository.GetCart()
}

func AddToCart(item model.CartItem) error {
	return repository.AddToCart(item)
}

func RemoveFromCart(id uint) error {
	return repository.RemoveFromCart(id)
}
