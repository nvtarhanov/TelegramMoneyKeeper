package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
	"github.com/nvtarhanov/TelegramMoneyKeeper/router"
)

// ReceiveMessage struct
type ReceiveMessage struct {
	UpdateID    int         `json:"update_id"`
	Message     Message     `json:"message"`
	ChannelPost ChannelPost `json:"channel_post"`
}

// Message struct
type Message struct {
	MessageID int        `json:"message_id"`
	From      From       `json:"from"`
	Chat      Chat       `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Entities  []Entities `json:"entities"`
}

// ChannelPost struct
type ChannelPost struct {
	MessageID int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

// // SendMessage struct
// type SendMessage struct {
// 	Ok     bool   `json:"ok"`
// 	Result Result `json:"result"`
// }

// // Result struct
// type Result struct {
// 	MessageID int    `json:"message_id"`
// 	Date      int    `json:"date"`
// 	Text      string `json:"text"`
// 	From      From   `json:"from"`
// 	Chat      Chat   `json:"chat"`
// }

// From struct
type From struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// Chat struct
type Chat struct {
	ID                          int    `json:"id"`
	FirstName                   string `json:"first_name"`
	UserName                    string `json:"username"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

// Entities struct
type Entities struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

const (
	TOKEN = "2141495170:AAHxcBtBSR2ALwk88DRBZfah0q_Td_cFFdc"
	URL   = "https://api.telegram.org/bot"
	PORT  = "80"
)

func main() {

	//1.Init config
	cfg, error := config.Init()

	if error != nil {
		log.Fatal(error)
	}

	//2.Init database

	if err := db.Init(cfg.DbConfig); err != nil {
		log.Fatal(err)
	}

	//Better to move into db package
	db.GetDB().AutoMigrate(&model.Account{}, &model.Entrie{}, &model.Transaction{})
	//fmt.Println(cfg.DbConfig)

	//3.Start router

	// http.HandleFunc("/api/v1/update", update)

	// fmt.Println("Listenning on port", PORT, ".")
	// if err := http.ListenAndServe(":"+PORT, nil); err != nil {
	// 	log.Fatal(err)
	// }

	r := router.Init()

	r.Run(":80")

}

func update(w http.ResponseWriter, r *http.Request) {

	message := &ReceiveMessage{}

	// chatID := 0
	// msgText := ""

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Println(err)
	}

	// if private or group
	if message.Message.Chat.ID != 0 {
		fmt.Println(message.Message.Chat.ID, message.Message.Text)
		// chatID = message.Message.Chat.ID
		// msgText = message.Message.Text
	} else {
		// if channel
		fmt.Println(message.ChannelPost.Chat.ID, message.ChannelPost.Text)
		// chatID = message.ChannelPost.Chat.ID
		// msgText = message.ChannelPost.Text
	}

	// respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=Received: %s", URL, TOKEN, chatID, msgText)

	// // send echo resp
	// _, err = http.Get(respMsg)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
