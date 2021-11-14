package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
)

func CreateTransaction(account *model.Account, value int) error {

	transaction := model.Transaction{AccountID: account.ID, Value: value}

	if err := database.GetDB().Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
