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
)

func SwitchState(inState int, inCommand string) (string, int) {

	outputeMessage := ""
	outputState := WaitForCommand

	switch inCommand {
	case command.CommandStart:
		//Account registration
		//errorMessage = cs.RegisterAccount(userID)
		outputState = WaitForNameRegistration
		outputeMessage = message.WaitForName
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
		//return help
		outputState = WaitForCommand
	case command.CommandGetProfileData:
		//return profile data
		outputState = WaitForCommand
	case command.CommandGetCalculation:
		//return money calculation
		//outputeMessage = cs.GetCalculatedData(userID)
		outputState = WaitForCommand
	}

	return outputeMessage, outputState
}
