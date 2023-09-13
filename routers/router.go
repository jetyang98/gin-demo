package routers

import (
	"gin-demo/handlers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handlers.Hello)

	return r
}
