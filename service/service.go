package service

type CommandService interface {
	RegisterAccount(chatID int) string
	SetNameByID(chatID int, data string) string
	SetMoneyGoalByID(chatID int, data string) string
	SetStartSumByID(chatID int, data string) string
	// UpdateState()
	// GetState()
	SetSalaryPerMonth()
	SetOutcomePerMonth()
	GetCalculatedData()
}

type TransportService interface {
	GetState()
	UpdateState()
}
