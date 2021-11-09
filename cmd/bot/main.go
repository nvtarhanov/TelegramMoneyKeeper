package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
	"github.com/nvtarhanov/TelegramMoneyKeeper/router"
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

	//3.Setup webhook
	data := url.Values{
		"url": {cfg.Ngrok_url + "/api/v1/update"},
	}
	_, err := http.PostForm("https://api.telegram.org/bot"+cfg.Telegram_token+"/setWebhook", data)

	if err != nil {
		log.Fatal("Unable to setup webhook")
	}

	//4.Start router
	r := router.Init()

	r.Run(":" + cfg.Port)

}

// // if private or group
// if message.Message.Chat.ID != 0 {
// 	fmt.Println(message.Message.Chat.ID, message.Message.Text)
// 	// chatID = message.Message.Chat.ID
// 	// msgText = message.Message.Text
// } else {
// 	// if channel
// 	fmt.Println(message.ChannelPost.Chat.ID, message.ChannelPost.Text)
// 	// chatID = message.ChannelPost.Chat.ID
// 	// msgText = message.ChannelPost.Text
// }
