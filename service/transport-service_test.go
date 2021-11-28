package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/nvtarhanov/TelegramMoneyKeeper/repository/mocks"
	state "github.com/nvtarhanov/TelegramMoneyKeeper/service/stateMachine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	mockRepository *mock_repository.MockStateRepository
}

func (s *Suite) SetupSuite() {
	mockCtrl := gomock.NewController(s.T())
	s.mockRepository = mock_repository.NewMockStateRepository(mockCtrl)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_UpdateState() {

	type mockBehavior func(mr *mock_repository.MockStateRepository)

	testTAble := []struct {
		name           string
		inputChatID    int
		inputState     int
		mockBehavior   mockBehavior
		expectedOutput error
	}{
		{name: "Update state without error",
			inputChatID:    1,
			inputState:     1,
			expectedOutput: nil,
			mockBehavior: func(mr *mock_repository.MockStateRepository) {
				mr.EXPECT().UpdateState(1, 1).Return(nil).Times(1)
			}},
		{name: "Create state error = gorm.ErrRecordNotFound",
			inputChatID:    1,
			inputState:     1,
			expectedOutput: nil,
			mockBehavior: func(mr *mock_repository.MockStateRepository) {
				mr.EXPECT().UpdateState(1, 1).Return(gorm.ErrRecordNotFound).Times(1)
				mr.EXPECT().WriteState(1, 1).Return(nil).Times(1)
			}},
		{name: "Update state with error <> gorm.ErrRecordNotFound",
			inputChatID:    1,
			inputState:     1,
			expectedOutput: errors.New("Mocked error"),
			mockBehavior: func(mr *mock_repository.MockStateRepository) {
				mr.EXPECT().UpdateState(1, 1).Return(gorm.ErrRecordNotFound).Times(1)
				mr.EXPECT().WriteState(1, 1).Return(errors.New("Mocked error")).Times(1)
			},
		},
	}

	for _, testtestCase := range testTAble {
		s.T().Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(s.T())
			defer mockCtrl.Finish()
			mockStateRepository := mock_repository.NewMockStateRepository(mockCtrl)
			testtestCase.mockBehavior(mockStateRepository)
			transportServiceHandler := NewTransportServiceHandler(mockStateRepository)
			err := transportServiceHandler.UpdateState(testtestCase.inputChatID, testtestCase.inputState)
			assert.Equal(s.T(), err, testtestCase.expectedOutput)
		})
	}

}

func (s *Suite) Test_GetState() {

	type mockBehavior func(mr *mock_repository.MockStateRepository)

	testTable := []struct {
		name                string
		inputChatID         int
		expectedOutputState int
		expectedOutputError error
		mockBehavior        mockBehavior
	}{
		{
			name:                "Get state without error",
			inputChatID:         1,
			expectedOutputState: 1,
			expectedOutputError: nil,
			mockBehavior: func(mr *mock_repository.MockStateRepository) {
				mr.EXPECT().GetCurrentStateByID(1).Return(1, nil)
			},
		}, {
			name:                "Get state with error gorm.ErrRecordNotFound",
			inputChatID:         1,
			expectedOutputState: state.WaitForRegistration,
			expectedOutputError: nil,
			mockBehavior: func(mr *mock_repository.MockStateRepository) {
				mr.EXPECT().GetCurrentStateByID(1).Return(state.WaitForRegistration, gorm.ErrRecordNotFound)
			},
		}, {
			name:                "Get state with error <> gorm.ErrRecordNotFound",
			inputChatID:         1,
			expectedOutputState: state.Error,
			expectedOutputError: errors.New("Mocked error"),
			mockBehavior: func(mr *mock_repository.MockStateRepository) {
				mr.EXPECT().GetCurrentStateByID(1).Return(state.Error, errors.New("Mocked error"))
			},
		},
	}

	for _, testtestCase := range testTable {
		s.T().Run(testtestCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(s.T())
			defer mockCtrl.Finish()
			mockStateRepository := mock_repository.NewMockStateRepository(mockCtrl)
			testtestCase.mockBehavior(mockStateRepository)
			transportServiceHandler := NewTransportServiceHandler(mockStateRepository)
			state, err := transportServiceHandler.GetState(testtestCase.inputChatID)
			assert.Equal(s.T(), err, testtestCase.expectedOutputError)
			assert.Equal(s.T(), state, testtestCase.expectedOutputState)
		})
	}
}
