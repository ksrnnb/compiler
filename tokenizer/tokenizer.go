package tokenizer

import (
	"bufio"
	"errors"
	"fmt"
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

type KeywordType int

const (
	Class KeywordType = iota + 1
	Method
	Function
	Constructor
	Int
	Boolean
	Char
	Void
	Var
	Static
	Field
	Let
	Do
	If
	Else
	While
	Return
	True
	False
	Null
	This
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

func (t Tokenizer) Keyword() (KeywordType, error) {
	tokenType, err := t.TokenType()
	if err != nil {
		return 0, fmt.Errorf("Keyword(): token type should be keyword, %v", err)
	}

	if tokenType != Keyword {
		return 0, errors.New("Keyword(): token type should be keyword")
	}

	return keyword(t.currentToken)
}

func (t Tokenizer) Symbol() (string, error) {
	tokenType, err := t.TokenType()
	if err != nil {
		return "", fmt.Errorf("Symbol(): token type should be symbol, %v", err)
	}

	if tokenType != Symbol {
		return "", errors.New("Symbol(): token type should be symbol")
	}

	return t.currentToken, nil
}
