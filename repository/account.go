package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
)

//GORM realisation
func CreateAccount(chatId int) error {

	account := model.Account{ID: chatId}

	database := database.GetDB()

	result := database.Create(&account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetName(a *model.Account, name string) error {

	a.Name = name

	database := database.GetDB()

	result := database.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetMoneyGoal(a *model.Account, moneyGoal int) error {

	a.MoneyGoal = moneyGoal

	database := database.GetDB()

	result := database.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetStartSum(a *model.Account, startsum int) error {

	a.Startsum = startsum

	database := database.GetDB()

	result := database.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAccountBySessionID(chatId int) (*model.Account, error) {

	account := model.Account{ID: chatId}

	database := database.GetDB()

	result := database.First(&account, "id = ?", account.ID)

	if result.Error != nil {
		return &account, result.Error
	}

	return &account, nil

}
