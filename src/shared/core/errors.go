package shared_core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ThrowUnauthorized() (int, any) {
	return http.StatusUnauthorized, gin.H{
		"code":    "UNAUTHORIZED",
		"message": "this api need authorization",
	}
}

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
func ThrowErrorWhenCreatingValueObject(err error) (int, any) {
	return http.StatusBadRequest, gin.H{
		"code":    "BAD_REQUEST",
		"message": fmt.Sprintf("occurred error when creating ValueObject : %s", err),
	}
}
