package assembler

type tokenType string

type token struct {
	Type    tokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	SYMBOL    = "SYMBOL"
	INT       = "INT"
	AT        = "@"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	SLASH     = "/"
	EQUAL     = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	AMPERSAND = "&"
	PIPE      = "|"

	COMMENT_SINGLE    = "//"
	COMMENT_MULTILINE = "/* */"

	// Keywords
	A      = "A"
	M      = "M"
	D      = "D"
	SP     = "SP"
	LCL    = "LCL"
	ARG    = "ARG"
	THIS   = "THIS"
	THAT   = "THAT"
	R0     = "R0"
	R1     = "R1"
	R2     = "R2"
	R3     = "R3"
	R4     = "R4"
	R5     = "R5"
	R6     = "R6"
	R7     = "R7"
	R8     = "R8"
	R9     = "R9"
	R10    = "R10"
	R11    = "R11"
	R12    = "R12"
	R13    = "R13"
	R14    = "R14"
	R15    = "R15"
	SCREEN = "SCREEN"
	KBD    = "KBD"

	// Jumps
	JEQ = "JEQ"
	JGT = "JGT"
	JGE = "JGE"
	JLT = "JLT"
	JNE = "JNE"
	JLE = "JLE"
	JMP = "JMP"
)

func newToken(tt tokenType, literal string) token {
	return token{Type: tt, Literal: literal}
}

var reserved = map[string]tokenType{
	"JEQ":  JEQ,
	"JGT":  JGT,
	"JGE":  JGE,
	"JLT":  JLT,
	"JNE":  JNE,
	"JLE":  JLE,
	"JMP":  JMP,
	A:      "A",
	M:      "M",
	D:      "D",
	SP:     "SP",
	LCL:    "LCL",
	ARG:    "ARG",
	THIS:   "THIS",
	THAT:   "THAT",
	R0:     "R0",
	R1:     "R1",
	R2:     "R2",
	R3:     "R3",
	R4:     "R4",
	R5:     "R5",
	R6:     "R6",
	R7:     "R7",
	R8:     "R8",
	R9:     "R9",
	R10:    "R10",
	R11:    "R11",
	R12:    "R12",
	R13:    "R13",
	R14:    "R14",
	R15:    "R15",
	SCREEN: "SCREEN",
	KBD:    "KBD",
}

func lookupSymbol(symbol string) tokenType {
	if tok, ok := reserved[symbol]; ok {
		return tok
	}
	return SYMBOL
}
