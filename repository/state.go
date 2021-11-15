package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type StateRepository struct {
	*gorm.DB
}

func NewStateRepository(db *gorm.DB) *StateRepository {
	return &StateRepository{db}
}

func (sr *StateRepository) GetCurrentStateByID(chatID int) (int, error) {
	var us int

	currentState := model.State{ID: chatID}

	result := sr.First(&currentState, "ID = ?", currentState.ID)

	if result.Error != nil {
		return 0, result.Error
	}

	us = currentState.State

	return us, nil
}

func (sr *StateRepository) WriteState(chatID int, state int) error {
	currentState := model.State{ID: chatID, State: state}

	result := sr.Create(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (sr *StateRepository) UpdateState(chatID int, state int) error {

	currentState := model.State{ID: chatID, State: state}

	result := sr.Save(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
