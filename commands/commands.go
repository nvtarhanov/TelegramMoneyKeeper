package commands

func CommadsList() []string {

	var commadsList []string

	commadsList = append(commadsList, "/start")
	commadsList = append(commadsList, "/help")
	commadsList = append(commadsList, "/setgoal")
	commadsList = append(commadsList, "/setsum")
	commadsList = append(commadsList, "/setname")

	return commadsList
}

func IsCommand(command string) bool {

	commandList := CommadsList()

	for _, com := range commandList {
		if command == com {
			return true
		}
	}

	return false
}
