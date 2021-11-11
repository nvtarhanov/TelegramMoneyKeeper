package model

import "time"

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AccountID int
	Account   Account `gorm:"foreignKey:AccountID"`
	Value     int
}

func CreateTransaction(account *Account, value int) error {

	// transaction := Transaction{AccountID: account.ID, Value: value}

	return nil
}
