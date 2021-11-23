package model

import (
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AccountID int
	Account   Account `gorm:"foreignKey:AccountID"`
	Value     int
}
