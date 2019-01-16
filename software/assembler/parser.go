package assembler

import (
	"fmt"
)

type symbolTable struct {
	nextPointer int
	symbols     map[string]int
}

func (st symbolTable) hasSymbol(s string) bool {
	_, seen := st.symbols[s]
	return seen
}

func (st *symbolTable) addSymbol(s string) {
	st.symbols[s] = st.nextPointer
	st.nextPointer += 1
}

func newSymbolTable() symbolTable {
	st := symbolTable{nextPointer: 16, symbols: make(map[string]int)}

	st.symbols["SP"] = 0
	st.symbols["LCL"] = 1
	st.symbols["ARG"] = 2
	st.symbols["THIS"] = 3
	st.symbols["THAT"] = 4
	st.symbols["R0"] = 0
	st.symbols["R1"] = 1
	st.symbols["R2"] = 2
	st.symbols["R3"] = 3
	st.symbols["R4"] = 4
	st.symbols["R5"] = 5
	st.symbols["R6"] = 6
	st.symbols["R7"] = 7
	st.symbols["R8"] = 8
	st.symbols["R9"] = 9
	st.symbols["R10"] = 10
	st.symbols["R11"] = 11
	st.symbols["R12"] = 12
	st.symbols["R13"] = 13
	st.symbols["R14"] = 14
	st.symbols["R15"] = 15
	st.symbols["SCREEN"] = 16384
	st.symbols["KBD"] = 24576

	return st
}

type parser struct {
	l      *lexer
	errors []string

	lastToken token
	curToken  token
	peekToken token

	table symbolTable
}

func newParser(l *lexer) *parser {
	p := &parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	p.setupSymbolTable()

	return p
}

func (p *parser) reset() {
	p.l.position = 0
	p.l.readPosition = 0
	p.l.ch = p.l.input[0]

	p.nextToken()
	p.nextToken()
}

func (p *parser) setupSymbolTable() {
	p.table = newSymbolTable()

	for !p.curTokenIs(EOF) {
		if p.curTokenIs(LPAREN) {
			p.nextToken()
			if !p.table.hasSymbol(p.curToken.Literal) {
				p.table.addSymbol(p.curToken.Literal)
			} else {
				p.errors = append(p.errors, "error: you must only declare a label once")
			}
		}
		p.nextToken()
	}

	fmt.Printf("%+v\n", p)
	p.reset()
}

func (p *parser) parseProgram() [][]byte {
	code := [][]byte{}

	for !p.curTokenIs(EOF) {
		instruction := p.parseInstruction()
		if instruction != nil {
			code = append(code, instruction)
		}
		p.nextToken()
	}

	return code
}

func (p *parser) parseInstruction() []byte {
	switch p.curToken.Type {
	// Comments
	case COMMENT_SINGLE:
	case COMMENT_MULTILINE:
		return nil

	// A instructions
	case SYMBOL:
		//return p.parseAInstruction()
		return []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0}

	// C Instructions
	case A:
	case M:
	case D:
		return []byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}
	}
	return nil
}

func (p *parser) nextToken() {
	p.lastToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

	if p.curTokenIs(ILLEGAL) {
		msg := fmt.Sprintf("got illegal token: %s", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		p.nextToken()
	}
}

func (p *parser) curTokenIs(tt tokenType) bool {
	return p.curToken.Type == tt
}

func (p *parser) peekTokenIs(tt tokenType) bool {
	return p.peekToken.Type == tt
}

func (p *parser) expectPeek(tt tokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	}
	p.peekError(tt)
	return false
}

func (p *parser) peekError(tt tokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tt, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
