package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/spf13/viper"
)

type ReceiveMessage struct {
	UpdateID    int         `json:"update_id"`
	Message     Message     `json:"message"`
	ChannelPost ChannelPost `json:"channel_post"`
}

type Message struct {
	MessageID int        `json:"message_id"`
	From      From       `json:"from"`
	Chat      Chat       `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Entities  []Entities `json:"entities"`
}

type ChannelPost struct {
	MessageID int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

type From struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID                          int    `json:"id"`
	FirstName                   string `json:"first_name"`
	UserName                    string `json:"username"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

type Entities struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

func Handle(c *gin.Context) {

	var message ReceiveMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatID := message.Message.Chat.ID
	msgText := message.Message.Text

	fmt.Println(chatID, msgText, viper.GetInt("port"))

	if err := switchCommand(chatID, msgText); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func switchCommand(chatID int, msgText string) error {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal("Failed to get config data")
	}

	switch msgText {
	case "/start":
		sendMessage(cfg, chatID, "Fuck off start")
	case "/help":
		sendMessage(cfg, chatID, "Fuck off help")
	default:
		sendMessage(cfg, chatID, "Неопознанная команда")
	}

	return nil
}

//Send message - msgText into tg chat with id - chatID
func sendMessage(cfg *config.Config, chatID int, msgText string) {
	message := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", cfg.Telegram_url, cfg.Telegram_token, chatID, msgText)

	if _, err := http.Get(message); err != nil {
		log.Fatal(err)
	}
}
