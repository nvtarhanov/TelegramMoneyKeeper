package repository

import (
	"errors"
	"regexp"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/suite"
)

type testState struct {
	id    int
	state int
}

func (s *Suite) Test_GetCurrentStateByID() {

	testData := testState{id: 1, state: 2}
	query := regexp.QuoteMeta(`SELECT * FROM "states" WHERE ID = $1 AND "states"."id" = $2 ORDER BY "states"."id" LIMIT 1`)

	rows := s.mock.NewRows([]string{"id", "state"}).AddRow(testData.id, testData.state)

	s.mock.ExpectQuery(query).WithArgs(testData.id, testData.id).WillReturnRows(rows)

	state, err := s.StateRepository.GetCurrentStateByID(testData.id)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), state)
	assert.Equal(s.T(), state, testData.state)
}

func (s *Suite) Test_GetCurrentStateByIDError() {

	testData := testState{id: 1, state: 2}
	query := regexp.QuoteMeta(`SELECT * FROM "states" WHERE ID = $1 AND "states"."id" = $2 ORDER BY "states"."id" LIMIT 1`)

	s.mock.ExpectQuery(query).WithArgs(testData.id, testData.id).WillReturnError(errors.New("Some error"))

	state, err := s.StateRepository.GetCurrentStateByID(testData.id)

	assert.Error(s.T(), err)
	assert.NotNil(s.T(), state)
	assert.Equal(s.T(), state, 0)
}

func (s *Suite) Test_WriteState() {

	testData := testState{id: 1, state: 1}

	query := regexp.QuoteMeta(`INSERT INTO "states" ("state","id") VALUES ($1,$2) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).WithArgs(testData.id, testData.id).WillReturnRows(s.mock.NewRows([]string{"id"}).AddRow("1"))

	s.mock.ExpectCommit()

	err := s.StateRepository.WriteState(testData.id, testData.state)

	assert.NoError(s.T(), err)

}

func (s *Suite) Test_WriteStateError() {

	testData := testState{id: 1, state: 1}

	query := regexp.QuoteMeta(`INSERT INTO "states" ("state","id") VALUES ($1,$2) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).WithArgs(testData.id, testData.id).WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.StateRepository.WriteState(testData.id, testData.state)

	assert.Error(s.T(), err)

}

func (s *Suite) Test_UpdateState() {

	testData := testState{id: 1, state: 1}

	query := regexp.QuoteMeta(`INSERT INTO "states" ("state","id") VALUES ($1,$2) RETURNING "id"`)

	s.mock.ExpectBegin()

	s.mock.ExpectQuery(query).WithArgs(testData.id, testData.id).WillReturnError(errors.New("Some error"))

	s.mock.ExpectRollback()

	err := s.StateRepository.WriteState(testData.id, testData.state)

	assert.Error(s.T(), err)

}
