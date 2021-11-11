package model

import (
	"testing"

	"github.com/nvtarhanov/TelegramMoneyKeeper/util"
)

func TestCreateEntrie(t *testing.T) {
	chatId := util.RandomChatId()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	if err := CreateEntrie(chatId); err != nil {
		t.Errorf("Cannot create entrie")
	}
}

func TestGetEntrieByID(t *testing.T) {
	chatID := util.RandomChatId()

	if err := CreateAccount(chatID); err != nil {
		t.Errorf("Cannot create account")
	}

	if err := CreateEntrie(chatID); err != nil {
		t.Errorf("Cannot create entrie")
	}

	entrie, err := GetEntrieByAccountID(chatID)

	if err != nil {
		t.Errorf("Cannot find entrie by id")
	}

	if entrie.AccountID != chatID {
		t.Errorf("CIncorrect entrie id")
	}

}

func TestSetSalaryPerMonth(t *testing.T) {

	chatId := util.RandomChatId()
	salaryPerMonth := util.RandomSum()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	if err := CreateEntrie(chatId); err != nil {
		t.Errorf("Cannot create entrie")
	}

	entrie, err := GetEntrieByAccountID(chatId)

	if err != nil {
		t.Errorf("Cannot find entrie by id")
	}

	if err := SetSalaryPerMonth(entrie, salaryPerMonth); err != nil {
		t.Errorf("Cannot set salary per month")
	}

	if entrie.SalaryPerMonth != salaryPerMonth {
		t.Errorf("Invalid salary per month")
	}

}

func TestSetOutcomePerMonth(t *testing.T) {

	chatId := util.RandomChatId()
	outcomePerMonth := util.RandomSum()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	if err := CreateEntrie(chatId); err != nil {
		t.Errorf("Cannot create entrie")
	}

	entrie, err := GetEntrieByAccountID(chatId)

	if err != nil {
		t.Errorf("Cannot find entrie by id")
	}

	if err := SetOutcomePerMonth(entrie, outcomePerMonth); err != nil {
		t.Errorf("Cannot set outcome per month")
	}

	if entrie.OutcomePerMonth != outcomePerMonth {
		t.Errorf("Invalid outcome per month")
	}

}
