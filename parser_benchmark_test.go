package benchmark

import (
	"fmt"
	"strings"
	"testing"
	"architecture-lab-4/parser"
)

func BenchmarkParser(b *testing.B) {
	baseInputCommands := "print sdas\nsha1 234\nprint 233\n"
	currentInputCommands := baseInputCommands
	for i := 0; i < 20; i++ {
		currentInputCommands += currentInputCommands
		b.Run(fmt.Sprintf("NumberOfCommands=%d", strings.Count(currentInputCommands, "\n")), func(b *testing.B) {
			_ = parser.Parse(currentInputCommands)
		})
	}
}
