package repository

//Better to switch on https://github.com/hashicorp/go-memdb

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type StateRepositoryGorm struct {
	*gorm.DB
}

func NewStateRepository(db *gorm.DB) *StateRepositoryGorm {
	return &StateRepositoryGorm{db}
}

func (sr *StateRepositoryGorm) GetCurrentStateByID(chatID int) (int, error) {

	currentState := model.State{ID: chatID}

	result := sr.First(&currentState, "ID = ?", currentState.ID)

	if result.Error != nil {
		return 0, result.Error
	}

	return currentState.State, nil
}

func (sr *StateRepositoryGorm) WriteState(chatID int, state int) error {
	currentState := model.State{ID: chatID, State: state}

	result := sr.Create(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (sr *StateRepositoryGorm) UpdateState(chatID int, state int) error {

	currentState := model.State{ID: chatID, State: state}

	result := sr.Save(&currentState)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
