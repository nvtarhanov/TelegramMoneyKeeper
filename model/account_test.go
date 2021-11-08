package model

import (
	"testing"

	"github.com/nvtarhanov/TelegramMoneyKeeper/util"
)

func TestCreateAccount(t *testing.T) {
	chatId := util.RandomChatId()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

}

func TestGetAccountById(t *testing.T) {
	chatId := util.RandomChatId()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	account, err := GetAccountBySessionID(chatId)

	if err != nil {
		t.Errorf("Cannot find account by id")
	}

	if account.ChatId != chatId {
		t.Errorf("Incorrect chatId")
	}
}

func TestSetName(t *testing.T) {

	chatId := util.RandomChatId()
	testName := util.RandomName()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("Cannot find account by id")
	}

	if err := SetName(account, testName); err != nil {
		t.Errorf("Cannot set name for account")
	}

}

func TestSetMoneyGoal(t *testing.T) {

	chatId := util.RandomChatId()
	testMoneyGoal := util.RandomSumGoal()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("Cannot find account by id")
	}

	if err := SetMoneyGoal(account, testMoneyGoal); err != nil {
		t.Errorf("Cannot set money goal for account")
	}

}

func TestSetStartSum(t *testing.T) {

	chatId := util.RandomChatId()
	testStartSum := util.RandomStartSum()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("Cannot find account by id")
	}

	if err := SetStartSum(account, testStartSum); err != nil {
		t.Errorf("Cannot set start sum for account")
	}

}
