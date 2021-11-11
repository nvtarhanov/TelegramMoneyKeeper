package model

import (
	"time"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AccountID int
	Account   Account `gorm:"foreignKey:AccountID"`
	Value     int
}

func CreateTransaction(account *Account, value int) error {

	transaction := Transaction{AccountID: account.ID, Value: value}

	if err := db.GetDB().Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
