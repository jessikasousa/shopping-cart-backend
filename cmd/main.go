package main

import (
	"log"
	"shopping-cart-backend/internal/controller"
	"shopping-cart-backend/pkg/db"
	"github.com/gin-contrib/cors"

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/products", controller.SearchProductsController)
	router.GET("/cart", controller.GetCartController)
	router.POST("/cart", controller.AddToCartController)
	router.DELETE("/cart/:id", controller.RemoveFromCartController)

	router.Run(":8080")
}
