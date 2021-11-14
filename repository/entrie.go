package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
)

func CreateEntrie(chatID int) error {

	entrie := model.Entrie{AccountID: chatID}

	if err := database.GetDB().Create(&entrie).Error; err != nil {
		return err
	}

	return nil
}

func SetSalaryPerMonth(entrie *model.Entrie, value int) error {

	entrie.SalaryPerMonth = value

	if err := database.GetDB().Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func SetOutcomePerMonth(entrie *model.Entrie, value int) error {

	entrie.OutcomePerMonth = value

	if err := database.GetDB().Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func GetEntrieByAccountID(ChatID int) (*model.Entrie, error) {

	entrie := model.Entrie{AccountID: ChatID}

	if err := database.GetDB().First(&entrie, "account_id = ?", entrie.AccountID).Error; err != nil {
		return &entrie, err
	}

	return &entrie, nil
}
