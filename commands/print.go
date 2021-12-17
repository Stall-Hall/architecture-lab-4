package commands

import (
	"fmt"
	"architecture-lab-4/engine"
)

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop engine.Handler){
	fmt.Println(p.arg)
}
