package service

import (
	"errors"

	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
	state "github.com/nvtarhanov/TelegramMoneyKeeper/service/stateMachine"
	"gorm.io/gorm"
)

type TransportServiceHandler struct {
	repository.StateRepository
}

func NewTransportServiceHandler(repo repository.StateRepository) *TransportServiceHandler {
	return &TransportServiceHandler{repo}
}

func (ts *TransportServiceHandler) UpdateState(chatID int, state int) error {

	err := ts.StateRepository.UpdateState(chatID, state)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = ts.StateRepository.WriteState(chatID, state)
	}

	if err != nil {
		return err
	}

	return nil
}

func (ts *TransportServiceHandler) GetState(chatID int) (int, error) {

	currentState, err := ts.StateRepository.GetCurrentStateByID(chatID)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return state.WaitForRegistration, nil
	}

	if err != nil {
		return state.Error, err
	}

	return currentState, nil
}
