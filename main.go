package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/go-restaurant-management/middleware"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8900"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	//routes.FoodRoutes(router)
	//routes.InvoiceRoutes(router)
	//routes.MenuRoutes(router)
	//routes.TableRoutes(router)
	//routes.OrderRoutes(router)
	//routes.OrderItemRoutes(router)

	err := router.Run(":" + PORT)
	if err != nil {
		return
	}
}
