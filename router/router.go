package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/update", update)

	return r
}

func update(c *gin.Context) {
	fmt.Println("listening")
}
