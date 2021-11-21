package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type EntrieRepositoryGorm struct {
	*gorm.DB
}

func NewEntrieRepository(db *gorm.DB) *EntrieRepositoryGorm {
	return &EntrieRepositoryGorm{db}
}

func (er *EntrieRepositoryGorm) CreateEntrie(chatID int) error {

	entrie := model.Entrie{AccountID: chatID}

	if err := er.Create(&entrie).Error; err != nil {
		return err
	}

	return nil
}

func (er *EntrieRepositoryGorm) SetSalaryPerMonth(entrie *model.Entrie, value int) error {

	entrie.SalaryPerMonth = value

	if err := er.Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func (er *EntrieRepositoryGorm) SetOutcomePerMonth(entrie *model.Entrie, value int) error {

	entrie.OutcomePerMonth = value

	if err := er.Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func (er *EntrieRepositoryGorm) GetEntrieByAccountID(ChatID int) (*model.Entrie, error) {

	entrie := model.Entrie{AccountID: ChatID}
	//Get last by id Salary entrie
	if err := er.Last(&entrie, "account_id = ?", entrie.AccountID).Error; err != nil {
		return &entrie, err
	}

	return &entrie, nil
}
