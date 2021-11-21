package service

type CommandService interface {
	RegisterAccount(chatID int) string
	SetNameByID(chatID int, data string) string
	SetMoneyGoalByID(chatID int, data string) string
	SetStartSumByID(chatID int, data string) string
	SetSalaryPerMonth(chatID int, data string) string
	SetOutcomePerMonth(chatID int, data string) string
	GetCalculatedData(chatID int) string
	SetTransaction(chatID int, data string) string

	ProcessCommand(inState int, inCommand string, userID int) (string, int)
}

type TransportService interface {
	GetState(int) (int, error)
	UpdateState(int, int) error
}
