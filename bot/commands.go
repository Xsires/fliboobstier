package bot

type Command struct {
	Name        string
	Description string
}

type Commands struct {
	AdminCommands []Command
}

var botCommands = Commands{
	AdminCommands: []Command{
		Command{
			Name:        "list_admins",
			Description: "list_admins",
		},
		Command{
			Name:        "show_regex_action",
			Description: "show_regex_action",
		},

		Command{
			Name:        "list_regex_actions",
			Description: "list_regex_actions",
		},

		Command{
			Name:        "add_regex_action_element",
			Description: "add_regex_action_element",
		},
		Command{
			Name:        "remove_regex_action_element",
			Description: "remove_regex_action_element",
		},
	},
}

func (CommandsInstance *Commands) isAdminCommand(command string) bool {
	for _, adminCommand := range CommandsInstance.AdminCommands {
		if command == adminCommand.Name {
			return true
		}
	}
	return false
}

func (CommandsInstance *Commands) getAllCommands() []Command {
	var result []Command
	for _, adminCommand := range CommandsInstance.AdminCommands {
		result = append(result, adminCommand)
	}
	return result

}
