package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
	mock_repository "github.com/nvtarhanov/TelegramMoneyKeeper/repository/mocks"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/message"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAccount(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name            string
		inputChatID     int
		ExpectedMessage string
		mockBehavior    mockBehavior
	}{
		{
			name:            "Account exists",
			inputChatID:     1,
			ExpectedMessage: message.AccountExists,
			mockBehavior: func(ar *mock_repository.MockAccountRepository,
				tr *mock_repository.MockTransactionRepository,
				sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
			},
		}, {
			name:            "Create account error",
			inputChatID:     1,
			ExpectedMessage: message.CannotCreateAccount,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
				ar.EXPECT().CreateAccount(1).Return(errors.New("Mocked error"))
			},
		}, {
			name:            "Create Salary record error",
			inputChatID:     1,
			ExpectedMessage: message.CannotCreateAccount,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
				ar.EXPECT().CreateAccount(1).Return(nil)
				sr.EXPECT().CreateEntrie(1).Return(errors.New("Mocked error"))
			},
		}, {
			name:            "Register account without error",
			inputChatID:     1,
			ExpectedMessage: "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
				ar.EXPECT().CreateAccount(1).Return(nil)
				sr.EXPECT().CreateEntrie(1).Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.RegisterAccount(testtestCase.inputChatID)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}

func TestSetNameByID(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name            string
		inputChatID     int
		inputName       string
		ExpectedMessage string
		mockBehavior    mockBehavior
	}{
		{
			name:            "Error getting account by session id",
			inputChatID:     1,
			inputName:       "TestName",
			ExpectedMessage: message.CannotFindAccountByID,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
			},
		}, {

			name:            "Error trying set name",
			inputChatID:     1,
			inputName:       "TestName",
			ExpectedMessage: message.CannotSetNameForAccount,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				ar.EXPECT().SetName(&model.Account{}, "TestName").Return(errors.New("Mocked error"))
			},
		}, {
			name:            "Set name without errors",
			inputChatID:     1,
			inputName:       "TestName",
			ExpectedMessage: "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				ar.EXPECT().SetName(&model.Account{}, "TestName").Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.SetNameByID(testtestCase.inputChatID, testtestCase.inputName)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}

func TestSetMoneyGoalByID(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name            string
		inputChatID     int
		inputMoneyGoal  string
		ExpectedMessage string
		mockBehavior    mockBehavior
	}{
		{
			name:            "Error converting inputMoneyGoal into int",
			inputChatID:     1,
			inputMoneyGoal:  "IncorrectData",
			ExpectedMessage: message.YouShouldEnterANumber,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
			},
		}, {
			name:            "Error getting account by session id",
			inputChatID:     1,
			inputMoneyGoal:  "2000",
			ExpectedMessage: message.CannotFindAccountByID,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
			},
		}, {
			name:            "Error setting money goal",
			inputChatID:     1,
			inputMoneyGoal:  "2000",
			ExpectedMessage: message.CannotSetMoneyGoalForAccount,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				ar.EXPECT().SetMoneyGoal(&model.Account{}, 2000).Return(errors.New("Mocked error"))
			},
		}, {
			name:            "Set money goal without error",
			inputChatID:     1,
			inputMoneyGoal:  "2000",
			ExpectedMessage: "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				ar.EXPECT().SetMoneyGoal(&model.Account{}, 2000).Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.SetMoneyGoalByID(testtestCase.inputChatID, testtestCase.inputMoneyGoal)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}
}

func TestSetStartSumByID(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name            string
		inputChatID     int
		inputStartSum   string
		ExpectedMessage string
		mockBehavior    mockBehavior
	}{
		{
			name:            "Error converting inputStartSum into int",
			inputChatID:     1,
			inputStartSum:   "IncorrectData",
			ExpectedMessage: message.YouShouldEnterANumber,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
			},
		}, {
			name:            "Error getting account by session id",
			inputChatID:     1,
			inputStartSum:   "2000",
			ExpectedMessage: message.CannotFindAccountByID,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
			},
		}, {

			name:            "Error setting money goal",
			inputChatID:     1,
			inputStartSum:   "2000",
			ExpectedMessage: message.CannotSetStartSumForAccount,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				ar.EXPECT().SetStartSum(&model.Account{}, 2000).Return(errors.New("Mocked error"))
			},
		}, {

			name:            "Set start sum without error",
			inputChatID:     1,
			inputStartSum:   "2000",
			ExpectedMessage: "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				ar.EXPECT().SetStartSum(&model.Account{}, 2000).Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.SetStartSumByID(testtestCase.inputChatID, testtestCase.inputStartSum)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}

func TestSetSalaryPerMonth(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name                string
		inputChatID         int
		inputSalaryPerMonth string
		ExpectedMessage     string
		mockBehavior        mockBehavior
	}{
		{
			name:                "Error converting inputSalaryPerMonth into int",
			inputChatID:         1,
			inputSalaryPerMonth: "IncorrectData",
			ExpectedMessage:     message.YouShouldEnterANumber,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
			},
		}, {
			name:                "Error getting salary record by account id",
			inputChatID:         1,
			inputSalaryPerMonth: "2000",
			ExpectedMessage:     message.CannotSetSalaryPerMonth,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, errors.New("Mocked error"))
			},
		}, {

			name:                "Error setting salary per month",
			inputChatID:         1,
			inputSalaryPerMonth: "2000",
			ExpectedMessage:     message.CannotSetSalaryPerMonth,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, nil)
				sr.EXPECT().SetSalaryPerMonth(&model.Entrie{}, 2000).Return(errors.New("Mocked error"))
			},
		}, {

			name:                "Set salary per month without errors",
			inputChatID:         1,
			inputSalaryPerMonth: "2000",
			ExpectedMessage:     "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, nil)
				sr.EXPECT().SetSalaryPerMonth(&model.Entrie{}, 2000).Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.SetSalaryPerMonth(testtestCase.inputChatID, testtestCase.inputSalaryPerMonth)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}

func TestSetOutcomePerMonth(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name                 string
		inputChatID          int
		inputOutcomePerMonth string
		ExpectedMessage      string
		mockBehavior         mockBehavior
	}{
		{
			name:                 "Error converting inputOutcomePerMonth into int",
			inputChatID:          1,
			inputOutcomePerMonth: "IncorrectData",
			ExpectedMessage:      message.YouShouldEnterANumber,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
			},
		}, {
			name:                 "Error getting salary record by account id",
			inputChatID:          1,
			inputOutcomePerMonth: "2000",
			ExpectedMessage:      message.CannotSetOutcomePerMonth,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, errors.New("Mocked error"))
			},
		}, {
			name:                 "Error setting outcome per month",
			inputChatID:          1,
			inputOutcomePerMonth: "2000",
			ExpectedMessage:      message.CannotSetOutcomePerMonth,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, nil)
				sr.EXPECT().SetOutcomePerMonth(&model.Entrie{}, 2000).Return(errors.New("Mocked error"))
			},
		}, {
			name:                 "Set outcome per month without errors",
			inputChatID:          1,
			inputOutcomePerMonth: "2000",
			ExpectedMessage:      "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, nil)
				sr.EXPECT().SetOutcomePerMonth(&model.Entrie{}, 2000).Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.SetOutcomePerMonth(testtestCase.inputChatID, testtestCase.inputOutcomePerMonth)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}

func TestSetTransaction(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name                string
		inputChatID         int
		inputTransactionSum string
		ExpectedMessage     string
		mockBehavior        mockBehavior
	}{
		{
			name:                "Error converting inputTransactionSum into int",
			inputChatID:         1,
			inputTransactionSum: "IncorrectData",
			ExpectedMessage:     message.YouShouldEnterANumber,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
			},
		}, {
			name:                "Error getting account by session id",
			inputChatID:         1,
			inputTransactionSum: "2000",
			ExpectedMessage:     message.CannotFindAccountByID,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
			},
		}, {
			name:                "Error setting transaction",
			inputChatID:         1,
			inputTransactionSum: "2000",
			ExpectedMessage:     message.CannotSetTransaction,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				tr.EXPECT().CreateTransaction(&model.Account{}, 2000).Return(errors.New("Mocked error"))
			},
		}, {
			name:                "Creating transaction withou error",
			inputChatID:         1,
			inputTransactionSum: "2000",
			ExpectedMessage:     "",
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, nil)
				tr.EXPECT().CreateTransaction(&model.Account{}, 2000).Return(nil)
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.SetTransaction(testtestCase.inputChatID, testtestCase.inputTransactionSum)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}

func TestGetCalculatedData(t *testing.T) {

	type mockBehavior func(ar *mock_repository.MockAccountRepository,
		tr *mock_repository.MockTransactionRepository,
		sr *mock_repository.MockSalaryRecordRepository)

	testTable := []struct {
		name            string
		inputChatID     int
		ExpectedMessage string
		mockBehavior    mockBehavior
	}{
		{
			name:            "Error getting account by session id",
			inputChatID:     1,
			ExpectedMessage: message.CannotFindAccountByID,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{}, errors.New("Mocked error"))
			},
		}, {
			name:            "Error account.MoneyGoal == 0",
			inputChatID:     1,
			ExpectedMessage: message.ShouldSetupMoneyGoal,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{ID: 1, MoneyGoal: 0}, nil)
			},
		}, {
			name:            "Error account.StartSum == 0",
			inputChatID:     1,
			ExpectedMessage: message.ShouldSetupStartSum,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{ID: 1, MoneyGoal: 100, Startsum: 0}, nil)
			},
		}, {
			name:            "Error getting salary record by account id",
			inputChatID:     1,
			ExpectedMessage: message.CantFindSalaryRecord,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{ID: 1, MoneyGoal: 100, Startsum: 50}, nil)
				sr.EXPECT().GetEntrieByAccountID(1).Return(&model.Entrie{}, errors.New("Mocked error"))
			},
		}, {
			name:            "Error salaryRecord.OutcomePerMonth == 0",
			inputChatID:     1,
			ExpectedMessage: message.ShouldSetupOutcome,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{ID: 1, MoneyGoal: 100, Startsum: 50}, nil)
				sr.EXPECT().GetEntrieByAccountID(1).
					Return(&model.Entrie{ID: 1, SalaryPerMonth: 100, OutcomePerMonth: 0}, nil)
			},
		}, {
			name:            "Error getting transaction sum",
			inputChatID:     1,
			ExpectedMessage: message.CantCalculateMoneyTransactions,
			mockBehavior: func(ar *mock_repository.MockAccountRepository, tr *mock_repository.MockTransactionRepository, sr *mock_repository.MockSalaryRecordRepository) {
				ar.EXPECT().GetAccountBySessionID(1).Return(&model.Account{ID: 1, MoneyGoal: 100, Startsum: 50}, nil)
				sr.EXPECT().GetEntrieByAccountID(1).
					Return(&model.Entrie{ID: 1, SalaryPerMonth: 100, OutcomePerMonth: 50}, nil)
				tr.EXPECT().GetTransactionSum(1).Return(0, errors.New("Mocked error"))
			},
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockAccountRepository := mock_repository.NewMockAccountRepository(mockCtrl)
			mockTransactionRepository := mock_repository.NewMockTransactionRepository(mockCtrl)
			mockSalaryRecordRepository := mock_repository.NewMockSalaryRecordRepository(mockCtrl)

			mockRepository := repository.Repository{
				AccountRepository:      mockAccountRepository,
				TransactionRepository:  mockTransactionRepository,
				SalaryRecordRepository: mockSalaryRecordRepository}

			testtestCase.mockBehavior(mockAccountRepository, mockTransactionRepository, mockSalaryRecordRepository)

			commandServiceHandler := NewCommandServiceHandler(mockRepository)

			message := commandServiceHandler.GetCalculatedData(testtestCase.inputChatID)

			assert.Equal(t, message, testtestCase.ExpectedMessage)

		})
	}

}
