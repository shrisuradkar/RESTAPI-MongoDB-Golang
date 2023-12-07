package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shrisuradkar/RestAPI-MonogoDB-Golang/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/user", controllers.GetUser())
	incomingRoutes.GET("/user/:id", controllers.GetUserByID())
	incomingRoutes.POST("/user", controllers.CreateUser())
	incomingRoutes.PUT("/user/:id", controllers.UpdateUser())
	incomingRoutes.DELETE("/user/:id", controllers.DeleteUser())
}
