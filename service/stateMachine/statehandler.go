package stateMachine

import (
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/command"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/message"
)

func ProcessCommand(inState State, inCommand string) (string, State) {

	outputeMessage := ""
	outputState := WaitForCommand

	if command.IsCommand(inCommand) {
		// If inCommand is a command we need to return a new state and message for state commands or return data for GET(stateless commands)
		switch inCommand {
		case command.CommandStart:
			//Account registration
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
			outputState = WaitForCommand
		}

	} else {
		//if inCommande is data we need to process our state with this data
		switch inState {
		case WaitForRegistration:
			//1.create account
			//2.next step - enter name
			//3.set state - WaitForName
		case WaitForCommand:
			outputeMessage = message.UregisteredCommand
			outputState = WaitForCommand
		case WaitForGoal:
			//Set Goal
			outputState = WaitForCommand
		case WaitForSum:
			//Set Sum
			outputState = WaitForCommand
		case WaitForName:
			//Set Name
			outputState = WaitForCommand
		case WaitForSalary:
			//Set Salary
			outputState = WaitForCommand
		case WaitForOutcome:
			//Set Outcome
			outputState = WaitForCommand
		case WaitForTransaction:
			//Set Transaction
			outputState = WaitForCommand
		//Commands for registration
		case WaitForNameRegistration:
			//set name
			outputeMessage = message.WaitForSum
			outputState = WaitForSumRegistration
		case WaitForSumRegistration:
			//set sum
			outputeMessage = message.WaitForGoal
			outputState = WaitForGoalRegistration
		case WaitForGoalRegistration:
			//set Goal
			outputState = WaitForSalaryRegistration
		case WaitForSalaryRegistration:
			//set Salary
			outputState = WaitForOutcomeRegistration
		case WaitForOutcomeRegistration:
			//set Outcome
			outputState = WaitForCommand
		}
	}

	return outputeMessage, outputState
}
