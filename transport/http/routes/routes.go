package routes

import (
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	var user *model.User
	incomingRoutes.POST("/createUser", service.CreateUser())
}
