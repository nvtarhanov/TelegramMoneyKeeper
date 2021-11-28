package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/command"
	stateMachine "github.com/nvtarhanov/TelegramMoneyKeeper/service/stateMachine"
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
	newState := stateMachine.WaitForCommand
	messageToUser := ""

	currentState, err := tg.transportService.GetState(chatID)

	if err != nil {
		log.Print(err)
	}

	if command.IsCommand(msgText) && !command.IsStateLessCommand(msgText) {
		messageToUser, newState = stateMachine.SwitchState(msgText)
	} else if command.IsStateLessCommand(msgText) {
		_, currentState := stateMachine.SwitchState(msgText)
		messageToUser, newState = tg.service.ProcessCommand(currentState, msgText, chatID)
	} else {
		//msgText is data
		messageToUser, newState = tg.service.ProcessCommand(currentState, msgText, chatID)
	}

	if messageToUser != "" {
		sendMessage(chatID, messageToUser)
	}

	err = tg.transportService.UpdateState(chatID, newState)

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
