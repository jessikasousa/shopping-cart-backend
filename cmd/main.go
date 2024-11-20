package main

import (
	"shopping-cart-backend/internal/controller"
	"shopping-cart-backend/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDatabase()

	router := gin.Default()

	router.GET("/products", controller.SearchProductsController)
	router.GET("/cart", controller.GetCartController)
	router.POST("/cart", controller.AddToCartController)
	router.DELETE("/cart/:id", controller.RemoveFromCartController)

	router.Run(":8080")
}
