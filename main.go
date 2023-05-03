package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"social-food-api/src"
)

func main() {
	fmt.Print("now", timestamppb.Now().Seconds)
	r := gin.New()

	// Global middleware
	// GIN_MODE=release로 하더라도 Logger 미들웨어는 gin.DefaultWriter에 로그를 기록합니다.
	// 기본값 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 미들웨어는 panic이 발생하면 500 에러를 씁니다.
	r.Use(gin.Recovery())

	main_module.SetupRoutes(r.Group("/"))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "NOT_FOUND",
			"message": "Page not found",
		})
	})

	err := r.Run(":9999")
	if err != nil {
		log.Panicln("Server can not running : ", err.Error())
	}
}
