package assembler

import (
	"fmt"
)

func Compile(source string) {
	l := newLexer(source)
	p := newParser(l)

	out := p.parseProgram()

	if len(p.errors) > 0 {
		fmt.Println(p.errors)
	} else {
		fmt.Println(out)
	}
}

func Compile2(source string) {
	l := newLexer(source)

	tok := l.NextToken()

	for tok.Type != EOF {
		fmt.Println(fmt.Sprintf("%v", tok))
		tok = l.NextToken()

		if tok.Type == ILLEGAL {
			fmt.Println(fmt.Sprintf("Illegal -- %v", tok))
			break
		}
	}

}
