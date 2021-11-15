package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/nvtarhanov/TelegramMoneyKeeper/controller"
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/router"
	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service"
)

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
	stateRepository := repository.NewStateRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	entrieRepository := repository.NewEntrieRepository(db)

	userService := service.UserService{AccountRepository: userRepository, StateRepository: stateRepository, TransactionRepository: transactionRepository, SalaryRecordRepository: entrieRepository}
	telegramController := controller.TelegramController{UserService: userService}

	//3.Setup webhook
	data := url.Values{
		"url": {cfg.Ngrok_url + "/api/v1/update"},
	}
	_, err = http.PostForm(cfg.Telegram_url+cfg.Telegram_token+"/setWebhook", data)

	if err != nil {
		log.Fatal("Unable to setup webhook")
	}

	//4.Start router
	r := router.Init(&telegramController)

	r.Run(":" + cfg.Port)

}
