package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/go-restaurant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.RouterGroup) {

	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
