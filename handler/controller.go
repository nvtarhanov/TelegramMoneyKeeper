package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service"
	"github.com/spf13/viper"
)

type TelegramHandeler struct {
	service          service.CommandService
	transportService service.TransportService
}

func NewTelegramHandler(service service.CommandService, transportService service.TransportService) *TelegramHandeler {
	return &TelegramHandeler{service: service, transportService: transportService}
}

func (tg *TelegramHandeler) Handle(c *gin.Context) {

	var message ReceiveMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatID := message.Message.Chat.ID
	msgText := message.Message.Text

	state, err := tg.transportService.GetState(chatID)

	if err != nil {
		log.Print(err)
	}

	messageReceive, state := tg.service.ProcessCommand(state, msgText, chatID)

	if messageReceive != "" {
		sendMessage(chatID, messageReceive)
	}

	err = tg.transportService.UpdateState(chatID, state)

	if err != nil {
		log.Print(err)
	}
}

//Send message - msgText into tg chat with id - chatID
func sendMessage(chatID int, msgText string) {
	message := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text= %s", viper.GetString("telegram_url"), viper.GetString("telegram_token"), chatID, msgText)

	if _, err := http.Get(message); err != nil {
		log.Print(err)
	}
}
