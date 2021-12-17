package commands

import (
	"fmt"
	"crypto/sha1"
	"architecture-lab-4/engine"
)

type ShaCommand struct {
	Arg string
}

func (p *ShaCommand) Execute(loop engine.Handler){
	s := fmt.Sprintf("%x", sha1.Sum([]byte(p.Arg)))
	loop.Post(&PrintCommand{Arg: s})
}
