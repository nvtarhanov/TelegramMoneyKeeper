//For gorm + postgres we cant use ExpectExec https://github.com/DATA-DOG/go-sqlmock/issues/118
//but it works ok with ExpectQuery

package repository

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository          Repository
	transportRepository TransportRepository
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

	s.repository = *NewRepository(s.DB)
	s.transportRepository = *NewTransportRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type testData struct {
	accountID int
	value     int
	id        string
}

func (s *Suite) Test_TransactionRepository_CreateTransaction() {

	testData := testData{accountID: 123456, value: 5500, id: "0"}

	account := &model.Account{ID: 123456}

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "transactions" ("created_at","updated_at","account_id","value") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(AnyTime{}, AnyTime{}, testData.accountID, testData.value).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testData.id))

	s.mock.ExpectCommit()

	err := s.repository.TransactionRepository.CreateTransaction(account, testData.value)

	require.NoError(s.T(), err)

}

func (s *Suite) Test_TransactionRepository_CreateTransactionError() {

	testData := testData{accountID: 123456, value: 5500, id: "0"}

	account := &model.Account{ID: 123456}

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "transactions" ("created_at","updated_at","account_id","value") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(AnyTime{}, AnyTime{}, testData.accountID, testData.value).
		WillReturnError(errors.New("Some error"))
	s.mock.ExpectRollback()

	err := s.repository.TransactionRepository.CreateTransaction(account, testData.value)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_TransactionRepository_GetTransactionSum() {

	testData := testData{accountID: 123456, value: 5500}

	query := regexp.QuoteMeta(`SELECT sum(value) as Total FROM "transactions" WHERE account_id = $1`)

	rows := s.mock.NewRows([]string{"account_id", "value"}).AddRow(123456, 2500)

	s.mock.ExpectQuery(query).WithArgs(testData.accountID).WillReturnRows(rows)

	sum, err := s.repository.TransactionRepository.GetTransactionSum(testData.accountID)

	fmt.Println(sum)

	assert.NotNil(s.T(), sum)
	assert.NoError(s.T(), err)
}

func (s *Suite) Test_TransactionRepository_GetTransactionSumError() {

	testData := testData{accountID: 123456, value: 5500}

	query := regexp.QuoteMeta(`SELECT sum(value) as Total FROM "transactions" WHERE account_id = $1`)

	s.mock.ExpectQuery(query).WithArgs(testData.accountID).WillReturnError(errors.New("Some error"))

	sum, err := s.repository.TransactionRepository.GetTransactionSum(testData.accountID)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), sum, 0)

}
