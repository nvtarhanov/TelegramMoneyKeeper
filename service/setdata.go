package service

import (
	"strconv"

	"github.com/nvtarhanov/TelegramMoneyKeeper/interfaces"
)

type UserService struct {
	interfaces.AccountRepository
}

//    /start command in telegram
func (us *UserService) RegisterAccount(chatID int) string {

	//Check if an account exists
	_, err := us.GetAccountBySessionID(chatID)
	if err == nil {
		return "Welcome back!"
	}
	//Create account
	if err := us.CreateAccount(chatID); err != nil {
		return "Cannot create account"
	}

	return ""

}

func (us *UserService) SetNameByID(chatID int, data string) string {

	account, err := us.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := us.SetName(account, data); err != nil {
		return "Cannot set name for account"
	}

	return ""

}

func (us *UserService) SetMoneyGoalByID(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		return "You should enter number"
	}

	account, err := us.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := us.SetMoneyGoal(account, value); err != nil {
		return "Cannot set money goal for account"
	}

	return ""

}

func (us *UserService) SetStartSumByID(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		return "You should enter number"
	}

	account, err := us.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := us.SetStartSum(account, value); err != nil {
		return "Cannot set start sum for account"
	}

	return ""

}
