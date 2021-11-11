package businesslogick

import (
	"log"
	"strings"

	"github.com/nvtarhanov/TelegramMoneyKeeper/commands"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/config"
	"github.com/nvtarhanov/TelegramMoneyKeeper/telegramapi"
)

func SwitchCommand(chatID int, msgText string) error {

	msgText = strings.ToLower(msgText)

	_, err := config.GetConfig()
	if err != nil {
		log.Fatal("Failed to get config data")
		return err
	}

	if commands.IsCommand(msgText) {
		if msgText == commands.CommandStart {
			RegisterAccount(chatID)
		} else {
			if err := model.UpdateState(chatID, strings.ToLower(msgText)); err != nil {
				log.Fatal("Cannot write state")
				return err
			}
		}
		return nil
	}

	userState, err := model.GetCurrentStateByID(chatID)

	if err != nil {
		log.Fatal("Cannot get user state")
		return err
	}

	switch strings.ToLower(userState) {
	case commands.CommandStart:
		RegisterAccount(chatID)
	case commands.CommandHelp:
		GetHelp(chatID)
	case commands.CommandSetGoal:
		SetMoneyGoal(chatID, msgText)
	case commands.CommandSetSum:
		SetStartSum(chatID, msgText)
	case commands.CommandSetName:
		SetName(chatID, msgText)
	default:
		telegramapi.SendMessage(chatID, "Unregistered command")
	}

	return nil
}
