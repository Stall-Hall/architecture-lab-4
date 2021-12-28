package parser

import (
	"architecture-lab-4/commands"
	"architecture-lab-4/engine"
	"strings"
)

func Parse(commandLine string) engine.Command {
	if len(commandLine) <= 0 {
		return (&commands.PrintCommand{Arg: "SYNTAX ERROR: empty line"})
	}

	fields := strings.Fields(commandLine)

	command := fields[0]
	if command == "print" {
		if len(fields) > 2 {
			return (&commands.PrintCommand{Arg: "SYNTAX ERROR: too many arguments for " + command})
		}
		return (&commands.PrintCommand{Arg: fields[1]})
	} else if command == "sha1" {
		if len(fields) > 2 {
			return (&commands.PrintCommand{Arg: "SYNTAX ERROR: too many arguments for " + command})
		}
		return (&commands.ShaCommand{Arg: fields[1]})
	}
	return (&commands.PrintCommand{Arg: "SYNTAX ERROR: unknown command"})
}
