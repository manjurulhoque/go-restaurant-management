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
		PORT = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	v1 := router.Group("/v1/api")
	routes.UserRoutes(v1)
	//router.Use(middleware.Authentication())

	routes.FoodRoutes(v1)
	routes.InvoiceRoutes(v1)
	routes.MenuRoutes(v1)
	routes.TableRoutes(v1)
	routes.OrderRoutes(v1)
	routes.OrderItemRoutes(v1)

	err := router.Run(":" + PORT)
	if err != nil {
		return
	}
}
