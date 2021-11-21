package service

import (
	"fmt"
	"strconv"

	"github.com/nvtarhanov/TelegramMoneyKeeper/model"
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

func (cs *CommandServiceHandler) ProcessCommand(inState int, inCommand string, userID int) (string, int) {

	outputeMessage := ""
	outputState := state.WaitForCommand
	errorMessage := ""

	//log.Printf("state %v inCommand %v userID %v", inState, inCommand, userID)

	if command.IsCommand(inCommand) {
		// If inCommand is a command we need to return a new state and message for state commands or return data for GET(stateless commands)
		switch inCommand {
		case command.CommandStart:
			//Account registration
			errorMessage = cs.RegisterAccount(userID)
			outputState = state.WaitForNameRegistration
			outputeMessage = message.WaitForName
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
			outputeMessage = cs.GetCalculatedData(userID)
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
			errorMessage = cs.SetMoneyGoalByID(userID, inCommand)
			outputeMessage = message.GoalSetted
			outputState = state.WaitForCommand
		case state.WaitForSum:
			//Set Sum
			errorMessage = cs.SetStartSumByID(userID, inCommand)
			outputeMessage = message.SumSetted
			outputState = state.WaitForCommand
		case state.WaitForName:
			//Set Name
			errorMessage = cs.SetNameByID(userID, inCommand)
			outputeMessage = message.NameSetted
			outputState = state.WaitForCommand
		case state.WaitForSalary:
			//Set Salary
			errorMessage = cs.SetSalaryPerMonth(userID, inCommand)
			outputeMessage = message.SalarySetted
			outputState = state.WaitForCommand
		case state.WaitForOutcome:
			//Set Outcome
			errorMessage = cs.SetOutcomePerMonth(userID, inCommand)
			outputeMessage = message.OutocmeSetted
			outputState = state.WaitForCommand
		case state.WaitForTransaction:
			//Set Transaction
			errorMessage = cs.SetTransaction(userID, inCommand)
			outputeMessage = message.TransactionSetted
			outputState = state.WaitForCommand
		//Commands for registration
		case state.WaitForNameRegistration:
			//set name
			errorMessage = cs.SetNameByID(userID, inCommand)
			outputeMessage = message.WaitForSum
			outputState = state.WaitForSumRegistration
		case state.WaitForSumRegistration:
			//set sum
			errorMessage = cs.SetStartSumByID(userID, inCommand)
			outputeMessage = message.WaitForGoal
			outputState = state.WaitForGoalRegistration
		case state.WaitForGoalRegistration:
			//set Goal
			errorMessage = cs.SetMoneyGoalByID(userID, inCommand)
			outputeMessage = message.WaitForSalary
			outputState = state.WaitForSalaryRegistration
		case state.WaitForSalaryRegistration:
			//set Salary
			errorMessage = cs.SetSalaryPerMonth(userID, inCommand)
			outputeMessage = message.WaitForOutcome
			outputState = state.WaitForOutcomeRegistration
		case state.WaitForOutcomeRegistration:
			//set Outcome
			errorMessage = cs.SetOutcomePerMonth(userID, inCommand)
			outputeMessage = message.RegistrationCompleted
			outputState = state.WaitForCommand
		}
	}

	if errorMessage != "" {
		outputeMessage = errorMessage
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

	if err := cs.SalaryRecordRepository.CreateEntrie(chatID); err != nil {
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

func (cs *CommandServiceHandler) SetSalaryPerMonth(chatID int, salary string) string {

	value, err := strconv.Atoi(salary)

	if err != nil {
		return "You should enter number"
	}

	salaryRecord, err := cs.SalaryRecordRepository.GetEntrieByAccountID(chatID)

	if err != nil {
		return "Cant set salary per month"
	}

	err = cs.SalaryRecordRepository.SetSalaryPerMonth(salaryRecord, value)

	if err != nil {
		return "Cant set salary per month"
	}

	return ""
}

func (cs *CommandServiceHandler) SetOutcomePerMonth(chatID int, outcome string) string {

	value, err := strconv.Atoi(outcome)

	if err != nil {
		return "You should enter number"
	}

	salaryRecord, err := cs.SalaryRecordRepository.GetEntrieByAccountID(chatID)

	if err != nil {
		return "Cant set outcome per month"
	}

	err = cs.SalaryRecordRepository.SetOutcomePerMonth(salaryRecord, value)

	if err != nil {
		return "Cant set outcome per month"
	}

	return ""
}

func (cs *CommandServiceHandler) SetTransaction(chatID int, sum string) string {

	value, err := strconv.Atoi(sum)

	if err != nil {
		return "You should enter number"
	}

	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	err = cs.TransactionRepository.CreateTransaction(account, value)

	if err != nil {
		return "Cannot set transaction"
	}

	return ""
}

func (cs *CommandServiceHandler) GetCalculatedData(chatID int) string {

	//get account data for calculation
	account, err := cs.AccountRepository.GetAccountBySessionID(chatID)
	if err != nil {
		return "Cannot find account by id"
	}

	if account.MoneyGoal == 0 {
		return "You should set up your money goal, type /setmoneygoal"
	}

	if account.Startsum == 0 {
		return "You should set up your Startsum, type /setstartsum"
	}

	//get salary data for calculation
	salaryRecord, err := cs.SalaryRecordRepository.GetEntrieByAccountID(chatID)

	if err != nil {
		return "Cant find salary record"
	}

	if salaryRecord.OutcomePerMonth == 0 {
		return "You should set up your salary, type /setsalary"
	}

	//get sum of all transactions if they exist
	transactionSum, err := cs.TransactionRepository.GetTransactionSum(chatID)

	if err != nil {
		return "Cant calculate money transactions"
	}

	return createMessageToUser(*account, *salaryRecord, transactionSum)
}

func createMessageToUser(account model.Account, salaryRecord model.Entrie, transactionSum int) string {

	messageToUser := fmt.Sprintf("Your start sum is: %v \n"+
		"Money goal is: %v \n "+
		"Salary is: %v \n "+
		"Outcome per month is: %v \n"+
		"Sum of transactions is: %v \n",
		account.Startsum, account.MoneyGoal, salaryRecord.SalaryPerMonth, salaryRecord.OutcomePerMonth, transactionSum)

	actualMoneyGoal := (account.MoneyGoal - account.Startsum) + transactionSum

	if actualMoneyGoal < 0 {
		messageToUser = messageToUser + fmt.Sprintf("You already get yor goal and have %v free money", actualMoneyGoal*-1)
	} else {
		freeMoneyPerMonth := salaryRecord.SalaryPerMonth - salaryRecord.OutcomePerMonth
		monthToGetGoal := float64(actualMoneyGoal) / float64(freeMoneyPerMonth)

		messageToUser = messageToUser + fmt.Sprintf("You need %v month to get your goal", monthToGetGoal)
	}

	return messageToUser
}
