package model

import (
	"testing"

	"github.com/nvtarhanov/TelegramMoneyKeeper/util"
)

func TestCreateTransaction(t *testing.T) {

	chatId := util.RandomChatId()
	value := util.RandomStartSum()

	if err := CreateAccount(chatId); err != nil {
		t.Errorf("Cannot create account")
	}

	account, err := GetAccountBySessionID(chatId)

	if err != nil {
		t.Errorf("Cannot find account by id")
	}

	if err := CreateTransaction(account, value); err != nil {
		t.Errorf("Cannot create transaction")
	}

}
