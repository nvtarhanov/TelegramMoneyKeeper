package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

//Send message - msgText into tg chat with id - chatID
func SendMessage(chatID int, msgText string) {
	message := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text= %s", viper.GetString("telegram_url"), viper.GetString("telegram_token"), chatID, msgText)

	if _, err := http.Get(message); err != nil {
		log.Fatal(err)
	}
}
