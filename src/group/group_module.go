package group

import (
	"github.com/gin-gonic/gin"
	"social-food-api/src/group/command/controller"
)

func SetupUserModule(r *gin.RouterGroup) {
	r.POST("/", group_command_controller.CreateGroup)
}
