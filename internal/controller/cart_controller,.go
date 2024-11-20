package controller

import (
	"net/http"
	"shopping-cart-backend/internal/model"
	"shopping-cart-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCartController(c *gin.Context) {
	cart, err := service.GetCart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func AddToCartController(c *gin.Context) {
	var item model.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := service.AddToCart(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart"})
}

func RemoveFromCartController(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := service.RemoveFromCart(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
