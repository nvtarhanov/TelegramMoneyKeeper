package model

import (
	"errors"
	"time"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint `gorm:"primaryKey"`
	ChatId    int  `gorm:"index"`
	CreatedAt time.Time
	Name      string `gorm:"index"`
	MoneyGoal int
	Startsum  int
}

func CreateAccount(chatId int) error {

	account := Account{ChatId: chatId}

	database := db.GetDB()

	result := database.First(&account, "chat_id = ?", account.ChatId)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("Account already exists")
	}

	result = database.Create(&account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetName(a *Account, name string) error {

	a.Name = name

	database := db.GetDB()

	result := database.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetMoneyGoal(a *Account, moneyGoal int) error {

	a.MoneyGoal = moneyGoal

	database := db.GetDB()

	result := database.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetStartSum(a *Account, startsum int) error {

	a.Startsum = startsum

	database := db.GetDB()

	result := database.Save(&a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAccountBySessionID(chatId int) (*Account, error) {

	account := Account{ChatId: chatId}

	database := db.GetDB()

	result := database.First(&account, "chat_id = ?", account.ChatId)

	if result.Error != nil {
		return &account, result.Error
	}

	return &account, nil

}
