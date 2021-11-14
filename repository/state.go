package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/infrastructure/database"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
)

func GetCurrentStateByID(chatID int) (int, error) {
	var us int

	currentState := model.State{ID: chatID}

	database := database.GetDB()

	result := database.First(&currentState, "ID = ?", currentState.ID)

	if result.Error != nil {
		return 0, result.Error
	}

	us = currentState.State

	return us, nil
}

func WriteState(chatID int, state int) error {
	currentState := model.State{ID: chatID, State: state}

	database := database.GetDB()
	result := database.Create(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateState(chatID int, state int) error {

	currentState := model.State{ID: chatID, State: state}

	database := database.GetDB()

	result := database.Save(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
