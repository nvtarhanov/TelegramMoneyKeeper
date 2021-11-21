package state

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
