package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nvtarhanov/TelegramMoneyKeeper/controllers/businesslogick"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/telegramapi"
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

	_, err := config.GetConfig()
	if err != nil {
		log.Fatal("Failed to get config data")
		return err
	}

	switch strings.ToLower(msgText) {
	case "/start":
		businesslogick.RegisterAccount(chatID)
	case "/help":
		businesslogick.GetHelp(chatID)
	case "/setgoal":
		businesslogick.SetMoneyGoal(chatID)
	case "/setsum":
		businesslogick.SetStartSum(chatID)
	case "/setname":
		businesslogick.SetName(chatID)
	default:
		telegramapi.SendMessage(chatID, "Unregistered command")
	}

	return nil
}
