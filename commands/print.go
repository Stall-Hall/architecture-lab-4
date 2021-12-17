package commands

import (
	"fmt"
	"architecture-lab-4/engine"
)

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(loop engine.Handler){
	fmt.Println(p.Arg)
}
