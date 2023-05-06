package shared_core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ThrowErrorWhenRequestBodyEmpty() (int, any) {
	return http.StatusBadRequest, gin.H{
		"code":    "BAD_REQUEST",
		"message": "request body was empty",
	}
}

func ThrowErrorWhenSomeRequestValueNotExist() (int, any) {
	return http.StatusBadRequest, gin.H{
		"code":    "BAD_REQUEST",
		"message": "some value was not requested",
	}
}

// ThrowErrorWhenCreatingValueObject : ValueObject 생성시 오류 났을 경우 Response
func ThrowErrorWhenCreatingValueObject(c *gin.Context, voErrors []error) {
	for _, err := range voErrors {
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"code":    "BAD_REQUEST",
				"message": fmt.Sprintf("occurred error when creating ValueObject : %s", err),
			})
			c.Abort()
			return
		}
	}
}
