package service

import "github.com/nvtarhanov/TelegramMoneyKeeper/repository"

type TransportServiceHandler struct {
	repository.StateRepository
}

func NewTransportServiceHandler(repo repository.StateRepository) *TransportServiceHandler {
	return &TransportServiceHandler{repo}
}

func (ts *TransportServiceHandler) UpdateState() {

}

func (ts *TransportServiceHandler) GetState() {

}
