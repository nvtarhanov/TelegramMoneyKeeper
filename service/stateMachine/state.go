package stateMachine

const (
	WaitForCommand State = iota
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

type State int
