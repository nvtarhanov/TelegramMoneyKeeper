package commands

const (
	CommandStart   string = "/start"
	CommandHelp    string = "/help"
	CommandSetGoal string = "/setgoal"
	CommandSetSum  string = "/setsum"
	CommandSetName string = "/setname"
)

func CommandsList() []string {

	var commadsList []string

	commadsList = append(commadsList, CommandStart)
	commadsList = append(commadsList, CommandHelp)
	commadsList = append(commadsList, CommandSetGoal)
	commadsList = append(commadsList, CommandSetSum)
	commadsList = append(commadsList, CommandSetName)

	return commadsList
}

func IsCommand(command string) bool {

	commandList := CommandsList()

	for _, com := range commandList {
		if command == com {
			return true
		}
	}

	return false
}
