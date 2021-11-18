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
)

type CommandHandeler struct {
	service          service.CommandService
	transportService service.TransportService
}

func NewCommandHandeler(service service.CommandService, transportService service.TransportService) *CommandHandeler {
	return &CommandHandeler{service: service, transportService: transportService}
}

const (
	chatID int = 123456
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
	transportRepository := repository.NewTransportRepository(db)
	repository := repository.NewRepository(db)

	transportService := service.NewTransportServiceHandler(transportRepository)
	service := service.NewCommandServiceHandler(*repository)

	commandHandeler := NewCommandHandeler(service, transportService)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		msgText, _ := reader.ReadString('\n')

		state, err := commandHandeler.transportService.GetState(chatID)

		log.Printf("Current state is %v message is %v", state, msgText)

		if err != nil {
			log.Fatal(err)
		}

		message, state := commandHandeler.service.ProcessCommand(state, msgText)

		log.Printf("State after is %v message to user is %v", state, message)

		//log.Print(message, state)

	}

}
