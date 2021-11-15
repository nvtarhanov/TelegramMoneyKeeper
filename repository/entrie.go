package repository

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"gorm.io/gorm"
)

type EntrieRepository struct {
	*gorm.DB
}

func NewEntrieRepository(db *gorm.DB) *EntrieRepository {
	return &EntrieRepository{db}
}

func (er *EntrieRepository) CreateEntrie(chatID int) error {

	entrie := model.Entrie{AccountID: chatID}

	if err := er.Create(&entrie).Error; err != nil {
		return err
	}

	return nil
}

func (er *EntrieRepository) SetSalaryPerMonth(entrie *model.Entrie, value int) error {

	entrie.SalaryPerMonth = value

	if err := er.Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func (er *EntrieRepository) SetOutcomePerMonth(entrie *model.Entrie, value int) error {

	entrie.OutcomePerMonth = value

	if err := er.Save(entrie).Error; err != nil {
		return err
	}

	return nil
}

func (er *EntrieRepository) GetEntrieByAccountID(ChatID int) (*model.Entrie, error) {

	entrie := model.Entrie{AccountID: ChatID}

	if err := er.First(&entrie, "account_id = ?", entrie.AccountID).Error; err != nil {
		return &entrie, err
	}

	return &entrie, nil
}
