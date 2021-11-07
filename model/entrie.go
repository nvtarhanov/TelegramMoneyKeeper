package model

import "time"

type Entrie struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountID       int
	Account         Account `gorm:"foreignKey:AccountID"`
	SalaryPerMonth  int
	OutcomePerMonth int
}
