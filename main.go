package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/go-restaurant-management/database"
	"github.com/manjurulhoque/go-restaurant-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8900"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	//router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)

	err := router.Run(":" + PORT)
	if err != nil {
		return
	}
}
