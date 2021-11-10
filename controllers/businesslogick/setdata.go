package businesslogick

import (
	"log"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/telegramapi"
)

func RegisterAccount(chatID int) {

	//Check if an account exists
	if _, err := model.GetAccountBySessionID(chatID); err == nil {
		telegramapi.SendMessage(chatID, "Welcome back!")
		return
	}
	//Create account
	if err := model.CreateAccount(chatID); err != nil {
		log.Fatal("Cannot create account")
	}

}

func SetMoneyGoal(chatID int) {

}

func SetStartSum(chatID int) {

}
