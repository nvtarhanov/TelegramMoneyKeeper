package state

import (
	"testing"

	"github.com/nvtarhanov/TelegramMoneyKeeper/service/command"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/message"
	"github.com/stretchr/testify/assert"
)

func TestSwitchState(t *testing.T) {

	testTable := []struct {
		name            string
		inputCommand    string
		ExpectedState   int
		ExpectedMessage string
	}{
		{
			name:            "Input CommandStart",
			inputCommand:    command.CommandStart,
			ExpectedState:   WaitForRegistration,
			ExpectedMessage: "",
		}, {
			name:            "Input CommandSetGoal",
			inputCommand:    command.CommandSetGoal,
			ExpectedState:   WaitForGoal,
			ExpectedMessage: message.WaitForGoal,
		}, {
			name:            "Input CommandSetSum",
			inputCommand:    command.CommandSetSum,
			ExpectedState:   WaitForSum,
			ExpectedMessage: message.WaitForSum,
		}, {
			name:            "Input CommandSetName",
			inputCommand:    command.CommandSetName,
			ExpectedState:   WaitForName,
			ExpectedMessage: message.WaitForName,
		}, {
			name:            "Input CommandSetSalary",
			inputCommand:    command.CommandSetSalary,
			ExpectedState:   WaitForSalary,
			ExpectedMessage: message.WaitForSalary,
		}, {
			name:            "Input CommandSetOutcome",
			inputCommand:    command.CommandSetOutcome,
			ExpectedState:   WaitForOutcome,
			ExpectedMessage: message.WaitForOutcome,
		}, {
			name:            "Input CommandSetTransaction",
			inputCommand:    command.CommandSetTransaction,
			ExpectedState:   WaitForTransaction,
			ExpectedMessage: message.WaitForTransaction,
		}, {
			name:            "Input CommandHelp",
			inputCommand:    command.CommandHelp,
			ExpectedState:   WaitForCommand,
			ExpectedMessage: "",
		}, {
			name:            "Input CommandGetProfileData",
			inputCommand:    command.CommandGetProfileData,
			ExpectedState:   WaitForCommand,
			ExpectedMessage: "",
		}, {
			name:            "Input CommandGetCalculation",
			inputCommand:    command.CommandGetCalculation,
			ExpectedState:   WaitForCalculation,
			ExpectedMessage: "",
		},
	}

	for _, testtestCase := range testTable {
		t.Run(testtestCase.name, func(t *testing.T) {
			outputMessage, outputState := SwitchState(testtestCase.inputCommand)
			assert.Equal(t, outputMessage, testtestCase.ExpectedMessage)
			assert.Equal(t, outputState, testtestCase.ExpectedState)
		})
	}
}
