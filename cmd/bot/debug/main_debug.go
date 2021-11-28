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
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/command"

	stateMachine "github.com/nvtarhanov/TelegramMoneyKeeper/service/stateMachine"
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

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		msgText, _ := reader.ReadString('\n')
		msgText = strings.TrimRight(msgText, "\r\n")

		message := ""

		currentState, err := transportService.GetState(userID)
		newState := stateMachine.WaitForCommand

		if err != nil {
			log.Print(err)
		}

		if command.IsCommand(msgText) && !command.IsStateLessCommand(msgText) {
			message, newState = stateMachine.SwitchState(currentState, msgText)
		} else if command.IsStateLessCommand(msgText) {
			_, currentState := stateMachine.SwitchState(currentState, msgText)
			message, newState = service.ProcessCommand(currentState, msgText, userID)
		} else {
			//msgText is data
			message, newState = service.ProcessCommand(currentState, msgText, userID)
		}

		err = transportService.UpdateState(userID, newState)

		if err != nil {
			log.Print(err)
		}

		//Send message to user
		log.Printf("message: %v", message)

	}

}
