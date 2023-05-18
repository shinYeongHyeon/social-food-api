package user

import (
	"github.com/gin-gonic/gin"
	"social-food-api/src/user/command/controller"
)

func SetupUserModule(r *gin.RouterGroup) {
	r.POST("/", user_command_controller.CreateUser)
}
