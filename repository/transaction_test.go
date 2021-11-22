package repository

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	transactionRepository TransactionRepository
	transaction           *model.Transaction
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(gorm.Dialector)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.transactionRepository = NewTransactionRepository(s.DB)
}
