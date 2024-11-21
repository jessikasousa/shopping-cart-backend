package main

import (
	"log"
	"shopping-cart-backend/internal/controller"
	"shopping-cart-backend/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	db.ConnectDatabase()

	router := gin.Default()

	router.GET("/products", controller.SearchProductsController)
	router.GET("/cart", controller.GetCartController)
	router.POST("/cart", controller.AddToCartController)
	router.DELETE("/cart/:id", controller.RemoveFromCartController)

	router.Run(":8080")
}
