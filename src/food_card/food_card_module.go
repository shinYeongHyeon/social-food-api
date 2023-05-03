package food_card

import (
	"github.com/gin-gonic/gin"
	"social-food-api/src/food_card/command/controller"
)

func SetupFoodCardModule(r *gin.RouterGroup) {
	r.POST("/", food_card_command_controller.CreateFoodCard)
}
