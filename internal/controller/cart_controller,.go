package controller

import (
	"net/http"
	"shopping-cart-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCartController(c *gin.Context) {
	cart, err := service.GetCart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar o carrinho"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func AddToCartController(c *gin.Context) {
	var req struct {
		ProductID string  `json:"product_id" binding:"required"`
		Title     string  `json:"title" binding:"required"`
		Quantity  int     `json:"quantity" binding:"required"`
		Price     float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := service.AddItemToCart(req.ProductID, req.Title, req.Quantity, req.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar item ao carrinho"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item adicionado ao carrinho com sucesso"})
}

func RemoveFromCartController(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = service.RemoveItemFromCart(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover item do carrinho"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removido com sucesso"})
}
