package model

import "time"

type Account struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `gorm:"index"`
	MoneyGoal int
	startsum  int
}
