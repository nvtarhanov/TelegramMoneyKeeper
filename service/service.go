package service

import "github.com/nvtarhanov/TelegramMoneyKeeper/repository"

type CommandService interface {
	RegisterAccount(chatID int) string
	SetNameByID(chatID int, data string) string
	SetMoneyGoalByID(chatID int, data string) string
	SetStartSumByID(chatID int, data string) string
	UpdateState()
	GetState()
	SetSalaryPerMonth()
	SetOutcomePerMonth()
	GetCalculatedData()
}

type Service struct {
	CommandService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
