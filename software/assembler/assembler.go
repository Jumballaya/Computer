package assembler

import "fmt"

func Compile(source string) {
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
