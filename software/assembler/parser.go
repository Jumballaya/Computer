package assembler

import (
	"fmt"
	"strconv"
)

type symbolTable struct {
	nextPointer int
	symbols     map[string]int
}

func (st symbolTable) hasSymbol(s string) bool {
	_, seen := st.symbols[s]
	return seen
}

func (st *symbolTable) addSymbol(s string) int {
	st.symbols[s] = st.nextPointer
	st.nextPointer += 1
	return st.nextPointer - 1
}

func newSymbolTable() *symbolTable {
	st := &symbolTable{nextPointer: 16, symbols: make(map[string]int)}

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

	table *symbolTable
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

func (p *parser) setupSymbolTable() {
	p.table = newSymbolTable()
	ct := p.curToken
	lt := p.lastToken
	pt := p.peekToken

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

	p.l.Reset()
	p.curToken = ct
	p.lastToken = lt
	p.peekToken = pt

	p.nextToken()
	p.nextToken()
}

func (p *parser) parseProgram() []int {
	code := []int{}

	for !p.curTokenIs(EOF) {
		instruction := p.parseInstruction()
		if instruction != nil {
			code = append(code, instruction[0])
		}
		p.nextToken()
	}

	for _, b := range code {
		fmt.Println(intToBinary(b))
	}

	return code
}

func (p *parser) parseInstruction() []int {
	switch p.curToken.Type {
	// Comments
	case COMMENT_SINGLE:
		return nil
	case COMMENT_MULTILINE:
		return nil

	// A instructions
	case SYMBOL:
		return p.parseAInstruction()

	// C Instructions
	case INT:
		return p.parseCInstruction()
	case A:
		return p.parseCInstruction()
	case M:
		return p.parseCInstruction()
	case D:
		return p.parseCInstruction()
	}
	return nil
}

func (p *parser) parseAInstruction() []int {
	// Check if the symbol is a number or a label
	// Return the byte of the address location
	location, ok := p.table.symbols[p.curToken.Literal]
	if !ok {
		location = p.table.addSymbol(p.curToken.Literal)
	}

	return []int{location}
}

func (p *parser) parseCInstruction() []int {
	tokens := []token{}
	for !p.curTokenIs(NEWLINE) {
		tokens = append(tokens, p.curToken)
		p.nextToken()
	}

	instr := Instruction{tokens: tokens}
	return []int{instr.Parse()}
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

func intToBinary(i int) string {
	n := int64(i)
	return fmt.Sprintf("%016s", strconv.FormatInt(n, 2))
}
