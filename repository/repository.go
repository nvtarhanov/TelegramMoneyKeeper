package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(chatId int) error
	SetName(a *model.Account, name string) error
	SetMoneyGoal(a *model.Account, moneyGoal int) error
	SetStartSum(a *model.Account, startSum int) error
	GetAccountBySessionID(chatId int) (*model.Account, error)
}

type SalaryRecordRepository interface {
	CreateEntrie(chatID int) error
	SetSalaryPerMonth(entrie *model.Entrie, value int) error
	SetOutcomePerMonth(entrie *model.Entrie, value int) error
	GetEntrieByAccountID(ChatID int) (*model.Entrie, error)
}

type TransactionRepository interface {
	CreateTransaction(account *model.Account, value int) error
	GetTransactionSum(ChatID int) (int, error)
}

type Repository struct {
	AccountRepository
	SalaryRecordRepository
	TransactionRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AccountRepository:      NewUserRepository(db),
		SalaryRecordRepository: NewEntrieRepository(db),
		TransactionRepository:  NewTransactionRepository(db),
	}
}

type StateRepository interface {
	GetCurrentStateByID(chatID int) (int, error)
	WriteState(chatID int, state int) error
	UpdateState(chatID int, state int) error
}

type TransportRepository struct {
	StateRepository
}

func NewTransportRepository(db *gorm.DB) *TransportRepository {
	return &TransportRepository{NewStateRepository(db)}
}
