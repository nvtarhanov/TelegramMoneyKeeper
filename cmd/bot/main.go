package main

import (
	"log"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
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

	//fmt.Println(cfg.DbConfig)

	//3.Start router

}
