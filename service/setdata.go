package service

import (
	"log"
	"strconv"
)

//    /start command in telegram
func RegisterAccount(chatID int) {

	//Check if an account exists
	_, err := model.GetAccountBySessionID(chatID)
	if err == nil {
		telegramapi.SendMessage(chatID, "Welcome back!")
		return
	}
	//Create account
	if err := model.CreateAccount(chatID); err != nil {
		log.Fatal("Cannot create account")
		return
	}

}

func SetName(chatID int, data string) {

	account, err := model.GetAccountBySessionID(chatID)
	if err != nil {
		log.Fatal("Cannot find account by id")
		return
	}

	if err := model.SetName(account, data); err != nil {
		log.Fatal("Cannot set name for account")
		return
	}

}

func SetMoneyGoal(chatID int, data string) {

	value, err := strconv.Atoi(data)

	if err != nil {
		log.Fatal("You should enter number")
		return
	}

	account, err := model.GetAccountBySessionID(chatID)
	if err != nil {
		log.Fatal("Cannot find account by id")
		return
	}

	if err := model.SetMoneyGoal(account, value); err != nil {
		log.Fatal("Cannot set money goal for account")
		return
	}

}

func SetStartSum(chatID int, data string) {

	value, err := strconv.Atoi(data)

	if err != nil {
		log.Fatal("You should enter number")
		return
	}

	account, err := GetAccountBySessionID(chatID)
	if err != nil {
		log.Fatal("Cannot find account by id")
		return
	}

	if err := model.SetStartSum(account, value); err != nil {
		log.Fatal("Cannot set start sum for account")
		return
	}

}
