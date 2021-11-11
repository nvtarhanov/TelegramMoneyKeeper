package model

import (
	"time"

	"github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"
)

type Entrie struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountID       int
	Account         Account `gorm:"foreignKey:AccountID"`
	SalaryPerMonth  int
	OutcomePerMonth int
}

func CreateEntrie(chatID int) error {

	entrie := Entrie{AccountID: chatID}

	if err := db.GetDB().Create(&entrie).Error; err != nil {
		return err
	}

	return nil
}

func SetSalaryPerMonth(entrie *Entrie, value int) error {

	entrie.SalaryPerMonth = value

	if err := db.GetDB().Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func SetOutcomePerMonth(entrie *Entrie, value int) error {

	entrie.OutcomePerMonth = value

	if err := db.GetDB().Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func GetEntrieByAccountID(ChatID int) (*Entrie, error) {

	entrie := Entrie{AccountID: ChatID}

	if err := db.GetDB().First(&entrie, "account_id = ?", entrie.AccountID).Error; err != nil {
		return &entrie, err
	}

	return &entrie, nil
}
