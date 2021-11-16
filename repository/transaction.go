package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type TransactionRepositoryGorm struct {
	*gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryGorm {
	return &TransactionRepositoryGorm{db}
}

func (tr *TransactionRepositoryGorm) CreateTransaction(account *model.Account, value int) error {

	transaction := model.Transaction{AccountID: account.ID, Value: value}

	if err := tr.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
