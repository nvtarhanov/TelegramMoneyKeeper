package model

import "time"

type Account struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `gorm:"index"`
	MoneyGoal int
	Startsum  int
}
