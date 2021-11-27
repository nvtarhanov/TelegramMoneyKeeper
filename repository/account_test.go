package repository

import (
	"errors"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

func (s *Suite) Test_CreateAccount() {

	testAccount := model.Account{ID: 1}

	query := regexp.QuoteMeta(`INSERT INTO "accounts" ("created_at","name","money_goal","startsum","id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).WithArgs(AnyTime{}, "", 0, 0, testAccount.ID).
		WillReturnRows(s.mock.NewRows([]string{"id"}).AddRow("1"))

	s.mock.ExpectCommit()

	err := s.repository.AccountRepository.CreateAccount(testAccount.ID)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_CreateAccountError() {

	testAccount := model.Account{ID: 1}

	query := regexp.QuoteMeta(`INSERT INTO "accounts" ("created_at","name","money_goal","startsum","id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).WithArgs(AnyTime{}, "", 0, 0, testAccount.ID).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.AccountRepository.CreateAccount(testAccount.ID)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_SetName() {

	testAccount := model.Account{ID: 1}
	name := "testName"

	query := `UPDATE "accounts"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, name, 0, 0, testAccount.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) //no inserted id, 1 affected row

	s.mock.ExpectCommit()

	err := s.repository.AccountRepository.SetName(&testAccount, name)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_SetNameError() {

	testAccount := model.Account{ID: 1}
	name := "testName"

	query := `UPDATE "accounts"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, name, 0, 0, testAccount.ID).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.AccountRepository.SetName(&testAccount, name)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_SetMoneyGoal() {

	testAccount := model.Account{ID: 1, Name: "testName"}
	moneyGoal := 1000

	query := `UPDATE "accounts"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, testAccount.Name, moneyGoal, 0, testAccount.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) //no inserted id, 1 affected row

	s.mock.ExpectCommit()

	err := s.repository.AccountRepository.SetMoneyGoal(&testAccount, moneyGoal)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_SetMoneyGoalError() {
	testAccount := model.Account{ID: 1, Name: "testName"}
	moneyGoal := 1000

	query := `UPDATE "accounts"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, testAccount.Name, moneyGoal, 0, testAccount.ID).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.AccountRepository.SetMoneyGoal(&testAccount, moneyGoal)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_SetStartSum() {

	testAccount := model.Account{ID: 1, Name: "testName", MoneyGoal: 1000}
	startSum := 2000

	query := `UPDATE "accounts"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, testAccount.Name, testAccount.MoneyGoal, startSum, testAccount.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) //no inserted id, 1 affected row

	s.mock.ExpectCommit()

	err := s.repository.AccountRepository.SetStartSum(&testAccount, startSum)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_SetStartSumError() {

	testAccount := model.Account{ID: 1, Name: "testName", MoneyGoal: 1000}
	startSum := 2000

	query := `UPDATE "accounts"`

	s.mock.ExpectBegin()

	s.mock.ExpectExec(query).
		WithArgs(AnyTime{}, testAccount.Name, testAccount.MoneyGoal, startSum, testAccount.ID).
		WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.repository.AccountRepository.SetStartSum(&testAccount, startSum)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_GetAccountBySessionID() {

	testAccount := model.Account{ID: 1, Name: "testName", MoneyGoal: 1000, Startsum: 2000}

	query := regexp.QuoteMeta(`SELECT * FROM "accounts"`)

	rows := s.mock.NewRows([]string{"id", "created_at", "name", "money_goal", "startsum"}).
		AddRow(testAccount.ID, time.Time{}, testAccount.Name, testAccount.MoneyGoal, testAccount.Startsum)

	s.mock.ExpectQuery(query).
		WithArgs(testAccount.ID, testAccount.ID).
		WillReturnRows(rows)

	account, err := s.repository.AccountRepository.GetAccountBySessionID(testAccount.ID)

	assert.NotNil(s.T(), account)
	assert.NoError(s.T(), err)

}

func (s *Suite) Test_GetAccountBySessionIDError() {

	testAccount := model.Account{ID: 1, Name: "testName", MoneyGoal: 1000, Startsum: 2000}

	query := regexp.QuoteMeta(`SELECT * FROM "accounts"`)

	s.mock.ExpectQuery(query).
		WithArgs(testAccount.ID, testAccount.ID).
		WillReturnError(errors.New("Some error"))

	account, err := s.repository.AccountRepository.GetAccountBySessionID(testAccount.ID)

	assert.NotNil(s.T(), account)
	assert.Error(s.T(), err)

}
