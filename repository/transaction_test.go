package repository

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	transactionRepository TransactionRepository
	//transaction           *model.Transaction
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, _ = sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})

	require.NoError(s.T(), err)

	s.transactionRepository = NewTransactionRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *Suite) Test_TransactionRepository_CreateTransaction() {
	var (
		accountID   = 123456
		value       = 5500
		currentDate = time.Now().Round(0)
		id          = "0"
	)

	account := &model.Account{ID: accountID}

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "transactions" ("created_at","updated_at","account_id","value") VALUES ($1,$2,$3,$4) RETURNING "id"`)).WithArgs(currentDate, currentDate, accountID, value).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(id))

	s.mock.ExpectCommit()

	err := s.transactionRepository.CreateTransaction(account, value)

	require.NoError(s.T(), err)

}
