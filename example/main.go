package main

import (
	"architecture-lab-4/engine"
	"architecture-lab-4/parser"
	"bufio"
	"fmt"
	"os"
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
			cmd := parser.Parse(nextCommandLine)
			eventLoop.Post(cmd)
		}
	} else {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	eventLoop.AwaitFinish()
}
