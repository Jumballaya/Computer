package assembler

import "bytes"

type lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *lexer) Reset() {
	l.position = 0
	l.readPosition = 0
	l.ch = l.input[0]
}

func newLexer(input string) *lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) NextToken() token {
	var tok token

	l.skipWhiteSpace()

	switch l.ch {
	case '@':
		l.readChar()
		literal := l.readSymbol()
		tok = token{Type: SYMBOL, Literal: literal}
	case 'M':
		tok = newToken(M, string(l.ch))
	case 'D':
		tok = newToken(D, string(l.ch))
	case 'A':
		tok = newToken(A, string(l.ch))
	case 'J':
		peek := l.peekChar()
		if peek == 'L' {
			l.readChar()
			if l.peekChar() == 'T' {
				l.readChar()
				tok = token{Type: JLT, Literal: "JLT"}
			}
			if l.peekChar() == 'E' {
				l.readChar()
				tok = token{Type: JLE, Literal: "JLE"}
			}
		}
		if peek == 'G' {
			l.readChar()
			if l.peekChar() == 'T' {
				l.readChar()
				tok = token{Type: JGT, Literal: "JGT"}
			}
			if l.peekChar() == 'E' {
				l.readChar()
				tok = token{Type: JGE, Literal: "JGE"}
			}
		}
		if peek == 'E' {
			l.readChar()
			if l.peekChar() == 'Q' {
				l.readChar()
				tok = token{Type: JEQ, Literal: "JEQ"}
			}
		}
		if peek == 'N' {
			if l.peekChar() == 'E' {
				l.readChar()
				tok = token{Type: JNE, Literal: "JNE"}
			}
		}
		if peek == 'M' {
			if l.peekChar() == 'P' {
				l.readChar()
				tok = token{Type: JMP, Literal: "JMP"}
			}
		}
	case '=':
		tok = token{Type: EQUAL, Literal: string(l.ch)}
	case '-':
		tok = token{Type: MINUS, Literal: string(l.ch)}
	case '+':
		tok = token{Type: PLUS, Literal: string(l.ch)}
	case ';':
		tok = token{Type: SEMICOLON, Literal: string(l.ch)}
	case '!':
		tok = token{Type: BANG, Literal: string(l.ch)}
	case '&':
		tok = token{Type: AMPERSAND, Literal: string(l.ch)}
	case '|':
		tok = token{Type: PIPE, Literal: string(l.ch)}
	case '(':
		tok = token{Type: LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token{Type: RPAREN, Literal: string(l.ch)}
	case '/':
		if l.peekChar() == '/' {
			literal := l.readComment(true)
			tok = token{Type: COMMENT_SINGLE, Literal: literal}
		}
		if l.peekChar() == '*' {
			// multiline comment
			literal := l.readComment(false)
			tok = token{Type: COMMENT_MULTILINE, Literal: literal}
		}
	case '\n':
		tok = token{Type: NEWLINE, Literal: "\\n"}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			literal := l.readSymbol()
			tok = token{Type: SYMBOL, Literal: literal}
		} else if isDigit(l.ch) {
			tok = token{Type: INT, Literal: string(l.ch)}
		} else {
			tok = newToken(ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *lexer) readComment(single bool) string {
	out := bytes.Buffer{}

	if single {
		for l.ch != '\n' {
			// DEAL WITH ERROR
			out.WriteByte(l.ch)
			l.readChar()
		}
	} else {
		for {
			if l.ch == '*' && l.peekChar() == '/' {
				out.WriteByte(l.ch)
				l.readChar()
				out.WriteByte(l.ch)
				l.readChar()
				break
			}
			out.WriteByte(l.ch)
			l.readChar()
		}
	}

	return out.String()
}

func (l *lexer) readString(initial byte) string {
	position := l.position + 1
	if initial == '\'' || initial == '"' {
		for {
			l.readChar()
			if l.ch == '\'' || l.ch == '"' || l.ch == 0 {
				break
			}
		}
		return l.input[position:l.position]
	} else {
		for {
			l.readChar()
			if l.ch == initial || l.ch == 0 {
				break
			}
		}
		return l.input[position:l.position]
	}
}

func (l *lexer) readSymbol() string {
	if l.ch == '@' {
		l.readChar()
	}
	pos := l.position
	l.readChar()
	for isSymbol(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '\''
}

func isSymbol(ch byte) bool {
	return isLetter(ch) || (ch != '.' && isDigit(ch))
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
