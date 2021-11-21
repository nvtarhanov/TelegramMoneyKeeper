package command

const (
	//Set commands
	CommandStart          string = "/start"
	CommandSetGoal        string = "/setgoal"
	CommandSetSum         string = "/setsum"
	CommandSetName        string = "/setname"
	CommandSetSalary      string = "/setsalary"
	CommandSetOutcome     string = "/setoutcome"
	CommandSetTransaction string = "/settransaction"

	//Get commands
	CommandHelp           string = "/help"
	CommandGetProfileData string = "/getprofiledata"
	CommandGetCalculation string = "/getcalculation"
)

func CommandsList() []string {

	var commadsList []string

	commadsList = append(commadsList, CommandStart)
	commadsList = append(commadsList, CommandHelp)
	commadsList = append(commadsList, CommandSetGoal)
	commadsList = append(commadsList, CommandSetSum)
	commadsList = append(commadsList, CommandSetName)
	commadsList = append(commadsList, CommandSetSalary)
	commadsList = append(commadsList, CommandSetOutcome)
	commadsList = append(commadsList, CommandSetTransaction)
	commadsList = append(commadsList, CommandHelp)
	commadsList = append(commadsList, CommandGetProfileData)
	commadsList = append(commadsList, CommandGetCalculation)

	return commadsList
}

func IsCommand(command string) bool {

	for _, com := range CommandsList() {
		if command == com {
			return true
		}
	}

	return false
}
