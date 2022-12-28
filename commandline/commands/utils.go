package commands

import "github.com/kc-workspace/go-lib/mapper"

func getName(commands, args []string, config mapper.Mapper) (string, []string) {
	var name = config.Mi("internal").So("command", DEFAULT)
	if len(args) > 0 {
		for _, cmd := range commands {
			if args[0] == cmd {
				name = cmd
				args = args[1:]
				break
			}
		}
	}

	return name, args
}
