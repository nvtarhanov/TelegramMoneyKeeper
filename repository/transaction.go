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

type resultSum struct {
	Total int
}

func (tr *TransactionRepositoryGorm) GetTransactionSum(chatID int) (int, error) {

	resultSum := resultSum{0}

	result := tr.Model(&model.Transaction{}).Select("sum(value) as Total").Where("account_id = ?", chatID).Find(&resultSum)

	if result.Error != nil {
		return 0, result.Error
	}

	return resultSum.Total, nil
}
