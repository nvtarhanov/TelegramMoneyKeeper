package service

import (
	"strconv"

	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
)

type CommandServiceHandler struct {
	AccountRepository      repository.AccountRepository
	SalaryRecordRepository repository.SalaryRecordRepository
	TransactionRepository  repository.TransactionRepository
	StateRepository        repository.StateRepository
}

func (cs *CommandServiceHandler) RegisterAccount(chatID int) string {

	//Check if an account exists
	_, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err == nil {
		return "Welcome back!"
	}
	//Create account
	if err := cs.AccountRepository.CreateAccount(chatID); err != nil {
		return "Cannot create account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetNameByID(chatID int, data string) string {

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := cs.AccountRepository.SetName(account, data); err != nil {
		return "Cannot set name for account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetMoneyGoalByID(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		return "You should enter number"
	}

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := cs.AccountRepository.SetMoneyGoal(account, value); err != nil {
		return "Cannot set money goal for account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetStartSumByID(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		return "You should enter number"
	}

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := cs.AccountRepository.SetStartSum(account, value); err != nil {
		return "Cannot set start sum for account"
	}

	return ""

}

func UpdateState() {

}

func GetState() {

}

func SetSalaryPerMonth() {

}
func SetOutcomePerMonth() {

}

func GetCalculatedData() {

}
