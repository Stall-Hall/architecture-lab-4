package main

import (
	"architecture-lab-4/commands"
	"architecture-lab-4/engine"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputFile = "input.txt"

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	input, err := os.Open(inputFile)
	if err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			nextCommandLine := scanner.Text()
			cmd := parse(nextCommandLine)
			eventLoop.Post(cmd)
		}
	} else {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	eventLoop.AwaitFinish()
}

func parse(commandLine string) engine.Command {
	if len(commandLine) <= 0 {
		return (&commands.PrintCommand{Arg: "SYNTAX ERROR: empty line"})
	}

	fields := strings.Fields(commandLine)

	if fields[0] == "print" {
		if len(fields) > 2 {
			return (&commands.PrintCommand{Arg: "SYNTAX ERROR: too many arguments"})
		}
		return (&commands.PrintCommand{Arg: fields[1]})
	} else if fields[0] == "sha1" {
		if len(fields) > 2 {
			return (&commands.PrintCommand{Arg: "SYNTAX ERROR: too many arguments"})
		}
		return (&commands.ShaCommand{Arg: fields[1]})
	}
	return (&commands.PrintCommand{Arg: "SYNTAX ERROR: unknown command"})
}
