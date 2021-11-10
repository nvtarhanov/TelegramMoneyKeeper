package businesslogick

import (
	"log"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/telegramapi"
)

//    /start command in telegram
func RegisterAccount(chatID int) {

	//Check if an account exists
	if _, err := model.GetAccountBySessionID(chatID); err == nil {
		telegramapi.SendMessage(chatID, "Welcome back!")
		return
	}
	//Create account
	if err := model.CreateAccount(chatID); err != nil {
		log.Fatal("Cannot create account")
		return
	}

	if err := model.WriteState(chatID, model.Name); err != nil {
		log.Fatal("Cannot write state")
		return
	}

	telegramapi.SendMessage(chatID, "Write your name!")

}

func SetMoneyGoal(chatID int) {

}

func SetStartSum(chatID int) {

}
