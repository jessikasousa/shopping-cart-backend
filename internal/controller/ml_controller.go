package controller

import (
	"net/http"
	"shopping-cart-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SearchProductsController(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro de consulta 'q' é obrigatório"})
		return
	}

	products, err := service.SearchProducts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar produtos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
