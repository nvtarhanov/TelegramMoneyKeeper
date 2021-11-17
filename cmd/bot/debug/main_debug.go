package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/stateMachine"
)

func main() {

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
	repository := repository.NewRepository(db)
	service := service.NewCommandServiceHandler(*repository)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')

		message, state := stateMachine.ProcessCommand(stateMachine.WaitForCommand, text)

		fmt.Println(message, state)

	}

}
