package main

import (
	"log"

	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/router"
	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
)

// import (
// 	"log"
// 	"net/http"
// 	"net/url"

// 	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
// 	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
// 	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
// 	"github.com/nvtarhanov/TelegramMoneyKeeper/router"
// )

func main() {

	//1.Init config
	cfg, error := config.NewConfig()

	if error != nil {
		log.Fatal(error)
	}

	//2.Init database

	db, err := database.Init(cfg.DbConfig)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db)

	// 	//Better to move into db package
	// 	db.GetDB().AutoMigrate(&model.Account{}, &model.Entrie{}, &model.Transaction{}, &model.State{})

	// 	//3.Setup webhook
	// 	data := url.Values{
	// 		"url": {cfg.Ngrok_url + "/api/v1/update"},
	// 	}
	// 	_, err := http.PostForm(cfg.Telegram_url+cfg.Telegram_token+"/setWebhook", data)

	// 	if err != nil {
	// 		log.Fatal("Unable to setup webhook")
	// 	}

	//4.Start router
	r := router.Init()

	r.Run(":" + cfg.Port)

}
