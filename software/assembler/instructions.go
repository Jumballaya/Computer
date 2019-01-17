package assembler

import (
	"fmt"
	"strconv"
)

type Instruction struct {
	tokens []token
}

func (i *Instruction) Parse() int {
	// C instructions always have 111 on the left side
	base := "111"

	// dest bits
	dest := i.parseDest()

	// 'a' bit, 0 = M, 1 = A
	a := i.parseABit()

	// Control bits for the ALU
	control := i.parseControlBits(a)

	// parse semicolon and JMP
	jump := i.parseJumpBits()

	bString := base + a + control + dest + jump
	n, _ := strconv.ParseInt(bString, 2, 64)

	return int(n)
}

func (i *Instruction) parseDest() string {
	dest := "000"
	switch i.tokens[0].Type {
	case A:
		switch i.tokens[1].Type {
		case M:
			if i.tokens[2].Type == D {
				dest = "111"
			} else {
				dest = "101"
			}
		case D:
			dest = "110"
		default:
			dest = "100"
		}

	case M:
		if i.tokens[1].Type == D {
			dest = "011"
		} else {
			dest = "001"
		}

	case D:
		dest = "010"
	}

	// If there isn't an equal sign the we are just sending null ("000")
	for _, t := range i.tokens {
		if t.Type == EQUAL {
			return dest
		}
	}
	return "000"
}

func (i *Instruction) parseABit() string {
	ts := i.tokens
	for n, t := range i.tokens {
		if t.Type == EQUAL {
			ts = i.tokens[n+1:]
		}
	}

	if ts[0].Type == M {
		return "1"
	}

	if ts[0].Type == BANG || ts[0].Type == MINUS {
		if ts[1].Type == M {
			return "1"
		}
	}

	if ts[0].Type == D {
		if len(ts) > 1 {
			tt := ts[1].Type
			if len(ts) > 2 {
				if tt == PLUS || tt == MINUS || tt == AMPERSAND || tt == PIPE {
					if ts[2].Type == M {
						return "1"
					}
				}
			}
		}
	}

	return "0"
}

func (i *Instruction) parseControlBits(a string) string {
	ts := i.tokens
	for n, t := range i.tokens {
		if t.Type == EQUAL {
			ts = i.tokens[n+1:]
		}
	}

	if a == "0" {
		switch ts[0].Type {
		case INT:
			if ts[0].Literal == "0" {
				return "101010"
			}
			if ts[0].Literal == "1" {
				return "111111"
			}

		case MINUS:
			switch ts[1].Type {
			case INT:
				if ts[1].Literal == "1" {
					return "111010"
				}
			case D:
				return "001111"
			case A:
				return "110011"
			}

		case BANG:
			switch ts[1].Type {
			case D:
				return "001101"
			case A:
				return "110001"
			}

		case D:
			if len(ts) > 1 {
				if ts[1].Type == PLUS {
					if ts[2].Literal == "1" {
						return "011111"
					}
					if ts[2].Type == A {
						return "000010"
					}
				}
				if ts[1].Type == MINUS {
					if ts[2].Literal == "1" {
						return "001110"
					}
					if ts[2].Type == A {
						return "010011"
					}
				}
				if ts[1].Type == AMPERSAND {
					if ts[2].Type == A {
						return "000000"
					}
				}
				if ts[1].Type == PIPE {
					if ts[2].Type == A {
						return "010101"
					}
				}
			} else {
				return "001100"
			}

		case A:
			if len(ts) > 1 {
				if ts[1].Type == PLUS {
					if ts[2].Literal == "1" {
						return "110111"
					}
				}
				if ts[1].Type == MINUS {
					if ts[2].Literal == "1" {
						return "110010"
					}
					if ts[2].Type == D {
						return "000111"
					}
				}
			} else {
				return "110000"
			}
		}
	}

	if a == "1" {
		switch ts[0].Type {
		case BANG:
			if ts[1].Type == M {
				return "110001"
			}

		case MINUS:
			if ts[1].Type == M {
				return "110011"
			}

		case M:
			if len(ts) > 1 {
				if ts[1].Type == PLUS {
					if ts[2].Literal == "1" {
						return "110111"
					}
				}
				if ts[1].Type == MINUS {
					if ts[2].Literal == "1" {
						return "110010"
					}
					if ts[2].Type == D {
						return "000111"
					}
				}
			} else {
				return "110000"
			}

		case D:
			if len(ts) > 1 {
				if ts[1].Type == PLUS {
					if ts[2].Literal == M {
						return "000010"
					}
				}
				if ts[1].Type == MINUS {
					if ts[2].Type == M {
						return "010011"
					}
				}
				if ts[1].Type == AMPERSAND {
					if ts[2].Type == M {
						return "000000"
					}
				}
				if ts[1].Type == PIPE {
					if ts[2].Type == M {
						return "010101"
					}
				}
			}
		}
	}

	return "000000"
}

func (i *Instruction) parseJumpBits() string {
	ts := i.tokens
	for n, t := range i.tokens {
		if t.Type == SEMICOLON {
			ts = i.tokens[n+1:]
		}
	}

	fmt.Println(ts[0])

	return "000"
}
