package service

import (
	"strconv"

	"github.com/nvtarhanov/TelegramMoneyKeeper/repository"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/command"
	"github.com/nvtarhanov/TelegramMoneyKeeper/service/message"
	state "github.com/nvtarhanov/TelegramMoneyKeeper/service/stateMachine"
)

type CommandServiceHandler struct {
	repository.Repository
}

func NewCommandServiceHandler(repo repository.Repository) *CommandServiceHandler {
	return &CommandServiceHandler{repo}
}

func (cs *CommandServiceHandler) ProcessCommand(inState int, inCommand string) (string, int) {

	outputeMessage := ""
	outputState := state.WaitForCommand

	if command.IsCommand(inCommand) {
		// If inCommand is a command we need to return a new state and message for state commands or return data for GET(stateless commands)
		switch inCommand {
		case command.CommandStart:
			//Account registration

		case command.CommandSetGoal:
			outputeMessage = message.WaitForGoal
			outputState = state.WaitForGoal
		case command.CommandSetSum:
			outputeMessage = message.WaitForSum
			outputState = state.WaitForSum
		case command.CommandSetName:
			outputeMessage = message.WaitForName
			outputState = state.WaitForName
		case command.CommandSetSalary:
			outputeMessage = message.WaitForSalary
			outputState = state.WaitForSalary
		case command.CommandSetOutcome:
			outputeMessage = message.WaitForOutcome
			outputState = state.WaitForOutcome
		case command.CommandSetTransaction:
			outputeMessage = message.WaitForTransaction
			outputState = state.WaitForTransaction
		case command.CommandHelp:
			//return help
			outputState = state.WaitForCommand
		case command.CommandGetProfileData:
			//return profile data
			outputState = state.WaitForCommand
		case command.CommandGetCalculation:
			//return money calculation
			outputState = state.WaitForCommand
		}

	} else {
		//if inCommande is data we need to process our state with this data
		switch inState {
		case state.WaitForRegistration:
			//1.create account
			//2.next step - enter name
			//3.set state - WaitForName
		case state.WaitForCommand:
			outputeMessage = message.UregisteredCommand
			outputState = state.WaitForCommand
		case state.WaitForGoal:
			//Set Goal
			outputState = state.WaitForCommand
		case state.WaitForSum:
			//Set Sum
			outputState = state.WaitForCommand
		case state.WaitForName:
			//Set Name
			outputState = state.WaitForCommand
		case state.WaitForSalary:
			//Set Salary
			outputState = state.WaitForCommand
		case state.WaitForOutcome:
			//Set Outcome
			outputState = state.WaitForCommand
		case state.WaitForTransaction:
			//Set Transaction
			outputState = state.WaitForCommand
		//Commands for registration
		case state.WaitForNameRegistration:
			//set name
			outputeMessage = message.WaitForSum
			outputState = state.WaitForSumRegistration
		case state.WaitForSumRegistration:
			//set sum
			outputeMessage = message.WaitForGoal
			outputState = state.WaitForGoalRegistration
		case state.WaitForGoalRegistration:
			//set Goal
			outputState = state.WaitForSalaryRegistration
		case state.WaitForSalaryRegistration:
			//set Salary
			outputState = state.WaitForOutcomeRegistration
		case state.WaitForOutcomeRegistration:
			//set Outcome
			outputState = state.WaitForCommand
		}
	}

	return outputeMessage, outputState
}

func (cs *CommandServiceHandler) RegisterAccount(chatID int) string {

	//Check if an account exists
	_, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err == nil {
		return "Welcome back!"
	}
	//Create account
	if err := cs.AccountRepository.CreateAccount(chatID); err != nil {
		return "Cannot create account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetNameByID(chatID int, data string) string {

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := cs.AccountRepository.SetName(account, data); err != nil {
		return "Cannot set name for account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetMoneyGoalByID(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		return "You should enter number"
	}

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := cs.AccountRepository.SetMoneyGoal(account, value); err != nil {
		return "Cannot set money goal for account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetStartSumByID(chatID int, data string) string {

	value, err := strconv.Atoi(data)

	if err != nil {
		return "You should enter number"
	}

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if err := cs.AccountRepository.SetStartSum(account, value); err != nil {
		return "Cannot set start sum for account"
	}

	return ""

}

func (cs *CommandServiceHandler) SetSalaryPerMonth() {

}
func (cs *CommandServiceHandler) SetOutcomePerMonth() {

}

func (cs *CommandServiceHandler) GetCalculatedData() {

}
