package repository

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

type testEntrie struct {
	id              int
	accountID       int
	salaryPerMont   int
	outcomePerMonth int
	chatID          int
}

func (s *Suite) Test_CreateEntrie() {

	testData := testEntrie{id: 1, accountID: 1, salaryPerMont: 1000, outcomePerMonth: 200, chatID: 1}

	query := regexp.QuoteMeta(`INSERT INTO "entries" ("created_at","updated_at","account_id","salary_per_month","outcome_per_month") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).
		WithArgs(AnyTime{}, AnyTime{}, testData.chatID, 0, 0).
		WillReturnRows(s.mock.NewRows([]string{"id"}).AddRow("1"))

	s.mock.ExpectCommit()

	err := s.repository.SalaryRecordRepository.CreateEntrie(testData.accountID)

	assert.Equal(s.T(), err, nil)

}

func (s *Suite) Test_CreateEntrieError() {

	testData := testEntrie{id: 1, accountID: 1, salaryPerMont: 1000, outcomePerMonth: 200, chatID: 1}

	query := regexp.QuoteMeta(`INSERT INTO "entries" ("created_at","updated_at","account_id","salary_per_month","outcome_per_month") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).
		WithArgs(AnyTime{}, AnyTime{}, testData.chatID, 0, 0).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.SalaryRecordRepository.CreateEntrie(testData.accountID)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_SetSalaryPerMonth() {

	testData := model.Entrie{ID: 1, AccountID: 2, SalaryPerMonth: 1000, OutcomePerMonth: 200}
	salary := 2000

	query := `UPDATE "entries"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, AnyTime{}, testData.AccountID, salary, testData.OutcomePerMonth, testData.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) //no inserted id, 1 affected row

	s.mock.ExpectCommit()

	err := s.repository.SalaryRecordRepository.SetSalaryPerMonth(&testData, salary)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_SetSalaryPerMonthError() {

	testData := model.Entrie{ID: 1, AccountID: 2, SalaryPerMonth: 1000, OutcomePerMonth: 200}
	salary := 2000

	query := `UPDATE "entries"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, AnyTime{}, testData.AccountID, salary, testData.OutcomePerMonth, testData.ID).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.SalaryRecordRepository.SetSalaryPerMonth(&testData, salary)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_SetOutcomePerMonth() {

	testData := model.Entrie{ID: 1, AccountID: 2, SalaryPerMonth: 1000, OutcomePerMonth: 200}
	outcome := 2000

	query := `UPDATE "entries"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, AnyTime{}, testData.AccountID, testData.SalaryPerMonth, outcome, testData.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	s.mock.ExpectCommit()

	err := s.repository.SalaryRecordRepository.SetOutcomePerMonth(&testData, outcome)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_SetOutcomePerMonthError() {

	testData := model.Entrie{ID: 1, AccountID: 2, SalaryPerMonth: 1000, OutcomePerMonth: 200}
	outcome := 2000

	query := `UPDATE "entries"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, AnyTime{}, testData.AccountID, testData.SalaryPerMonth, outcome, testData.ID).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.SalaryRecordRepository.SetOutcomePerMonth(&testData, outcome)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_SGetEntrieByAccountID() {

	testData := model.Entrie{ID: 1, AccountID: 1, SalaryPerMonth: 1000, OutcomePerMonth: 200}

	query := regexp.QuoteMeta(`SELECT * FROM "entries"`)

	s.mock.ExpectQuery(query).WithArgs(testData.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "account_id", "salary_per_month", "outcome_per_month"}).
			AddRow(testData.ID, testData.AccountID, testData.SalaryPerMonth, testData.OutcomePerMonth))

	salaryEntrie, err := s.repository.SalaryRecordRepository.GetEntrieByAccountID(testData.AccountID)

	assert.NotNil(s.T(), salaryEntrie)
	assert.NoError(s.T(), err)

}

func (s *Suite) Test_SGetEntrieByAccountIDError() {

	testData := model.Entrie{ID: 1, AccountID: 1, SalaryPerMonth: 1000, OutcomePerMonth: 200}

	query := regexp.QuoteMeta(`SELECT * FROM "entries"`)

	s.mock.ExpectQuery(query).WithArgs(testData.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "account_id", "salary_per_month", "outcome_per_month"}).
			AddRow(testData.ID, testData.AccountID, testData.SalaryPerMonth, testData.OutcomePerMonth)).
		WillReturnError(errors.New("Some error"))

	salaryEntrie, err := s.repository.SalaryRecordRepository.GetEntrieByAccountID(testData.AccountID)

	assert.NotNil(s.T(), salaryEntrie)
	assert.Error(s.T(), err)

}
