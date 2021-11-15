package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nvtarhanov/TelegramMoneyKeeper/controller"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/update", controller.Handle)

	return r
}
