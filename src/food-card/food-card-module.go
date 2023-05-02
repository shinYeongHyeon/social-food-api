package food_card

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupFoodCardModule(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "FOOD_CARD OK",
		})
	})
}
