package tokenizer

import (
	"bufio"
	"io"
)

type TokenType int

const (
	Keyword TokenType = iota + 1
	Symbol
	Identifier
	IntConst
	StringConst
)

type Tokenizer struct {
	scanner      *bufio.Scanner
	isDone       bool
	currentToken string
}

func NewTokenizer(input io.Reader) *Tokenizer {
	sc := bufio.NewScanner(input)
	return &Tokenizer{scanner: sc, isDone: false}
}

func (t Tokenizer) HasMoreTokens() bool {
	return t.isDone
}

func (t *Tokenizer) Advance() {

}

func (t Tokenizer) TokenType() (TokenType, error) {
	return tokenType(t.currentToken)
}
