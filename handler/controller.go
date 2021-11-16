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
	service *service.Service
}

func NewTelegramHandler(service *service.Service) *TelegramHandeler {
	return &TelegramHandeler{service: service}
}

func (tg *TelegramHandeler) Handle(c *gin.Context) {

	var message ReceiveMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatID := message.Message.Chat.ID
	msgText := message.Message.Text

	fmt.Println(chatID, msgText, viper.GetInt("port"))

	// if err := businesslogick.SwitchCommand(chatID, msgText); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	//message, state := service.ProcessCommand(service.WaitForCommand, msgText)

	//messenger.SendMessage(chatID, messageText)

}

//Send message - msgText into tg chat with id - chatID
func (TelegramHandeler) SendMessage(chatID int, msgText string) {
	message := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text= %s", viper.GetString("telegram_url"), viper.GetString("telegram_token"), chatID, msgText)

	if _, err := http.Get(message); err != nil {
		log.Fatal(err)
	}
}
