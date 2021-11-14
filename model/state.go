package model

type State struct {
	ID    int `gorm:"primaryKey"`
	State int
}
