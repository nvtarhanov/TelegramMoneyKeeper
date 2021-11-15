package service

import (
	"strconv"
)

// type UserReposytory struct {

// }

//    /start command in telegram
func RegisterAccount(chatID int) string {

	//Check if an account exists
	_, err := model.GetAccountBySessionID(chatID)
	if err == nil {
		//telegramapi.SendMessage(chatID, "Welcome back!")
		return "Welcome back!"
	}
	//Create account
	if err := model.CreateAccount(chatID); err != nil {
		//log.Fatal("Cannot create account")
		return "Cannot create account"
	}

}

func SetName(chatID int, data string) string {

	account, err := model.GetAccountBySessionID(chatID)
	if err != nil {
		//log.Fatal("Cannot find account by id")
		return "Cannot find account by id"
	}

	if err := model.SetName(account, data); err != nil {
		//log.Fatal("Cannot set name for account")
		return "Cannot set name for account"
	}

}

func SetMoneyGoal(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		//log.Fatal("You should enter number")
		return "You should enter number"
	}

	account, err := model.GetAccountBySessionID(chatID)
	if err != nil {
		//log.Fatal("Cannot find account by id")
		return "Cannot find account by id"
	}

	if err := model.SetMoneyGoal(account, value); err != nil {
		//log.Fatal("Cannot set money goal for account")
		return "Cannot set money goal for account"
	}

}

func SetStartSum(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		//log.Fatal("You should enter number")
		return "You should enter number"
	}

	account, err := GetAccountBySessionID(chatID)
	if err != nil {
		//log.Fatal("Cannot find account by id")
		return "Cannot find account by id"
	}

	if err := model.SetStartSum(account, value); err != nil {
		//log.Fatal("Cannot set start sum for account")
		return "Cannot set start sum for account"
	}

}
