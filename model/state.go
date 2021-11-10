package model

import "github.com/nvtarhanov/TelegramMoneyKeeper/pkg/db"

type UserState string

const (
	Name      UserState = "/setname"
	MoneyGoal UserState = "/setsum"
	StartSum  UserState = "/setname"
)

type State struct {
	ID    int `gorm:"primaryKey"`
	State UserState
}

func GetCurrentStateByID(chatID int) (UserState, error) {
	var us UserState

	currentState := State{ID: chatID}

	database := db.GetDB()

	result := database.First(&currentState, "ID = ?", currentState.ID)

	if result.Error != nil {
		return "", result.Error
	}

	us = currentState.State

	return us, nil
}

func WriteState(chatID int, state UserState) error {
	currentState := State{ID: chatID, State: state}

	database := db.GetDB()
	result := database.Create(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
