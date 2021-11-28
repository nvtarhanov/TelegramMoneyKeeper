package state

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/command"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/message"
)

const (
	WaitForCommand int = iota + 1
	WaitForGoal
	WaitForSum
	WaitForName
	WaitForSalary
	WaitForOutcome
	WaitForTransaction

	WaitForRegistration
	WaitForGoalRegistration
	WaitForSumRegistration
	WaitForNameRegistration
	WaitForSalaryRegistration
	WaitForOutcomeRegistration
	Error
	WaitForCalculation
)

func SwitchState(inCommand string) (string, int) {

	outputeMessage := ""
	outputState := WaitForCommand

	switch inCommand {
	case command.CommandStart:
		//Account registration(stateless command)
		outputState = WaitForRegistration
		outputeMessage = ""
	case command.CommandSetGoal:
		outputeMessage = message.WaitForGoal
		outputState = WaitForGoal
	case command.CommandSetSum:
		outputeMessage = message.WaitForSum
		outputState = WaitForSum
	case command.CommandSetName:
		outputeMessage = message.WaitForName
		outputState = WaitForName
	case command.CommandSetSalary:
		outputeMessage = message.WaitForSalary
		outputState = WaitForSalary
	case command.CommandSetOutcome:
		outputeMessage = message.WaitForOutcome
		outputState = WaitForOutcome
	case command.CommandSetTransaction:
		outputeMessage = message.WaitForTransaction
		outputState = WaitForTransaction
	case command.CommandHelp:
		//return help(stateless command)
		outputState = WaitForCommand
	case command.CommandGetProfileData:
		//return profile data(stateless command)
		outputState = WaitForCommand
	case command.CommandGetCalculation:
		//return money calculation with empty message(because it is stateless command)
		outputState = WaitForCalculation
	}

	return outputeMessage, outputState
}
