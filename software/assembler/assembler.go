package assembler

import (
	"fmt"
)

func Compile(source string) {
	l := newLexer(source)
	p := newParser(l)

	out := p.parseProgram()

	fmt.Printf("%v\n", out)
}
