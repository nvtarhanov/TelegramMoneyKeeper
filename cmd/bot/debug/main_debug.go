package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

func main() {

	userID := 123456

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

	//commandHandeler := NewCommandHandeler(service, transportService)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		msgText, _ := reader.ReadString('\n')
		msgText = strings.TrimRight(msgText, "\r\n")

		state, err := transportService.GetState(userID)

		log.Printf("Current state is %v message is %v", state, msgText)

		if err != nil {
			log.Print(err)
		}

		message, state := service.ProcessCommand(state, msgText, userID)

		err = transportService.UpdateState(userID, state)

		if err != nil {
			log.Print(err)
		}

		log.Printf("State after is %v message to user is %v", state, message)

	}

}
