package commands

import (
	"log"
	"strings"

	"github.com/nvtarhanov/TelegramMoneyKeeper/controllers/businesslogick"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/telegramapi"
)

const (
	CommandStart   string = "/start"
	CommandHelp    string = "/help"
	CommandSetGoal string = "/setgoal"
	CommandSetSum  string = "/setsum"
	CommandSetName string = "/setname"
)

func CommandsList() []string {

	var commadsList []string

	commadsList = append(commadsList, CommandStart)
	commadsList = append(commadsList, CommandHelp)
	commadsList = append(commadsList, CommandSetGoal)
	commadsList = append(commadsList, CommandSetSum)
	commadsList = append(commadsList, CommandSetName)

	return commadsList
}

func IsCommand(command string) bool {

	commandList := CommandsList()

	for _, com := range commandList {
		if command == com {
			return true
		}
	}

	return false
}

func SwitchCommand(chatID int, msgText string) error {

	_, err := config.GetConfig()
	if err != nil {
		log.Fatal("Failed to get config data")
		return err
	}

	if !IsCommand(strings.ToLower(msgText)) {
		userState, err := model.GetCurrentStateByID(chatID)

		if err != nil {
			log.Fatal("Cannot get user state")
			return err
		}

		msgText = string(userState)
	}

	//data    message data , not a command
	//command - starts with /

	switch strings.ToLower(msgText) {
	case "/start":
		businesslogick.RegisterAccount(chatID)
	case "/help":
		businesslogick.GetHelp(chatID)
	case "/setgoal":
		businesslogick.SetMoneyGoal(chatID)
	case "/setsum":
		businesslogick.SetStartSum(chatID)
	case "/setname":
		businesslogick.SetName(chatID)
	default:
		telegramapi.SendMessage(chatID, "Unregistered command")
	}

	return nil
}
