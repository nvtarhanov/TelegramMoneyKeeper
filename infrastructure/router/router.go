package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.Default()

	//r.POST("/api/v1/update", controllers.Handle)

	return r
}
