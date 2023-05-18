package shared_core

import (
	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header["Authorization"]
		if authorization == nil {
			c.Set("isAuthorized", false)
			c.Next()
			return
		}

		if authorization[0] == "" {
			c.Set(IsAuthorized, false)
			c.Next()
			return
		}

		c.Set(IsAuthorized, true)
		c.Set(Authorized, authorization[0])
		c.Next()
	}
}
