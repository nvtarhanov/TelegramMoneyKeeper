package model

import (
	"time"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
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

	result := database.Create(&account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SetName(name string) error {

	return nil
}

func SetMoneyGoal(moneyGoal int) error {

	return nil
}

func SetStartSum(startsum int) error {

	return nil
}

func GetAccountBySessionID(chatId int) (Account, error) {

	acc := Account{}

	return acc, nil
}
