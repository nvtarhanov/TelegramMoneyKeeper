package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/nvtarhanov/TelegramMoneyKeeper/handler"
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

	//Inject dependency
	transportRepository := repository.NewTransportRepository(db)
	repository := repository.NewRepository(db)

	serviceTransport := service.NewTransportServiceHandler(transportRepository)
	service := service.NewCommandServiceHandler(*repository)

	handler := handler.NewTelegramHandler(service, serviceTransport)

	//3.Setup webhook
	data := url.Values{
		"url": {cfg.Ngrok_url + "/api/v1/update"},
	}
	_, err = http.PostForm(cfg.Telegram_url+cfg.Telegram_token+"/setWebhook", data)

	if err != nil {
		log.Fatal("Unable to setup webhook")
	}

	//4.Start router
	router := router.Init(handler)

	router.Run(":" + cfg.Port)

}
