package controller

import (
	"log"
	"net/http"
	"shopping-cart-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SearchProductsController(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		log.Println("Erro: Parâmetro de consulta 'q' não foi fornecido")
		c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro de consulta 'q' é obrigatório"})
		return
	}

	products, err := service.SearchProducts(query)
	if err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar produtos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
