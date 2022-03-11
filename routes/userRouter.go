package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/go-restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/user/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
}
