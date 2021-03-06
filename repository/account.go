package repository

import (
	"log"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{db}
}

func (db *UserRepositoryGorm) CreateAccount(chatId int) error {

	account := model.Account{ID: chatId}

	result := db.Create(&account)

	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	}

	return nil
}

func (db *UserRepositoryGorm) SetName(a *model.Account, name string) error {

	a.Name = name

	result := db.Save(&a)

	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	}

	return nil
}

func (db *UserRepositoryGorm) SetMoneyGoal(a *model.Account, moneyGoal int) error {

	a.MoneyGoal = moneyGoal

	result := db.Save(&a)

	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	}

	return nil
}

func (db *UserRepositoryGorm) SetStartSum(a *model.Account, startsum int) error {

	a.Startsum = startsum

	result := db.Save(&a)

	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	}

	return nil
}

func (db *UserRepositoryGorm) GetAccountBySessionID(chatId int) (*model.Account, error) {

	account := model.Account{ID: chatId}

	result := db.First(&account, "id = ?", account.ID)

	if result.Error != nil {
		log.Print(result.Error)
		return &account, result.Error
	}

	return &account, nil

}
