package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (db *UserRepository) CreateAccount(chatId int) error {

	account := model.Account{ID: chatId}

	result := db.Create(&account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *UserRepository) SetName(a *model.Account, name string) error {

	a.Name = name

	result := db.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *UserRepository) SetMoneyGoal(a *model.Account, moneyGoal int) error {

	a.MoneyGoal = moneyGoal

	result := db.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *UserRepository) SetStartSum(a *model.Account, startsum int) error {

	a.Startsum = startsum

	result := db.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *UserRepository) GetAccountBySessionID(chatId int) (*model.Account, error) {

	account := model.Account{ID: chatId}

	result := db.First(&account, "id = ?", account.ID)

	if result.Error != nil {
		return &account, result.Error
	}

	return &account, nil

}
