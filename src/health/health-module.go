package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupHealthModule(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}
