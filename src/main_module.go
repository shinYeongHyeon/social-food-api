package main_module

import (
	"github.com/gin-gonic/gin"
	"social-food-api/src/food_card"
	"social-food-api/src/health"
)

func SetupRoutes(r *gin.RouterGroup) {
	health.SetupHealthModule(r.Group("/health"))
	food_card.SetupFoodCardModule(r.Group("/food-cards"))
}
