package commands

import (
	"fmt"
	"crypto/sha1"
	"architecture-lab-4/engine"
)

type shaCommand struct {
	arg string
}

func (p *shaCommand) Execute(loop engine.Handler){
	s := fmt.Sprintf("%x", sha1.Sum([]byte(p.arg)))
	loop.Post(&printCommand{arg: s})
}
