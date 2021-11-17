package state

const (
	WaitForCommand int = iota
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
)
