package REFACTORING

import (
	"testing"

	"github.com/nvtarhanov/TelegramMoneyKeeper/util"
)

func TestCreateAccount(t *testing.T) {
	chatId := util.RandomChatId()

	// Normal creation
	if err := CreateAccount(chatId); err != nil {
		t.Errorf("CreateAccount: %+v", err)
	}
	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("GetAccountBySessionID: %+v", err)
	}
	if account.ID != chatId {
		t.Errorf("CreateAccount error: account not found with ID %v", account.ID)
	}

	// Duplicate creation
	err = CreateAccount(chatId)
	if err == nil {
		t.Errorf("CreateAccount error: should not allow duplicate creation")
	}

	// SQL error test
	// Disable SQL mock and run CreateAccount(chatId) again, expecting an error
}

func TestGetAccountById(t *testing.T) {
	chatId := util.RandomChatId()
	if err := CreateAccount(chatId); err != nil {
		t.Errorf("CreateAccount: %+v", err)
	}

	// Normal search
	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("GetAccountBySessionID: %+v", err)
	}
	if account.ID != chatId {
		t.Errorf("GetAccountBySessionID error: searched for %v, got %v", chatId, account.ID)
	}

	// Search for non existing account
	_, err = GetAccountBySessionID(chatId + 1)
	if err == nil {
		t.Errorf("GetAccountBySessionID should return error for non existing account")
	}
	expected := "Some well-known error text" // I don`t know what it should be, find out by yourself
	if err.Error() != expected {
		t.Errorf("GetAccountBySessionID: expected %v error, got %v", expected, err)
	}

	// SQL error test
	// Disable SQL mock and run CreateAccount(chatId) again, expecting an error
}

func TestSetName(t *testing.T) {
	chatId := util.RandomChatId()
	testName := util.RandomName()
	if err := CreateAccount(chatId); err != nil {
		t.Errorf("CreateAccount: %+v", err)
	}
	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("GetAccountBySessionID: %+v", err)
	}

	// Normal operation
	if err := SetName(account, testName); err != nil {
		t.Errorf("SetName: %+v", err)
	}

	// SQL error test
	// Disable SQL mock and run CreateAccount(chatId) again, expecting an error
}

// Remaining tests are identical to TestSetName

func TestSetMoneyGoal(t *testing.T) {

	chatId := util.RandomChatId()
	testMoneyGoal := util.RandomSumGoal()

	// if err := CreateAccount(chatId); err != nil {
	// 	t.Errorf("Cannot create account")
	// }

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

	// if err := CreateAccount(chatId); err != nil {
	// 	t.Errorf("Cannot create account")
	// }

	account, err := GetAccountBySessionID(chatId)
	if err != nil {
		t.Errorf("Cannot find account by id")
	}

	if err := SetStartSum(account, testStartSum); err != nil {
		t.Errorf("Cannot set start sum for account")
	}

}
