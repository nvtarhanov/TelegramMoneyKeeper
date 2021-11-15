package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nvtarhanov/TelegramMoneyKeeper/controller"
)

func Init(tg *controller.TelegramController) *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/update", tg.Handle) //)controller.Handle)

	return r
}
