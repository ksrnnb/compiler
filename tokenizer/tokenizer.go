package tokenizer

import (
	"bufio"
	"io"

	"github.com/ksrnnb/compiler/utils"
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

func (t Tokenizer) TokenType() TokenType {
	if utils.IsInSlice(getKeywordTokens(), t.currentToken) {
		return Keyword
	}

	return 0
}

func getKeywordTokens() []string {
	return []string{
		"class",
		"constructor",
		"function",
		"method",
		"field",
		"static",
		"var",
		"int",
		"char",
		"boolean",
		"void",
		"true",
		"false",
		"null",
		"this",
		"let",
		"do",
		"if",
		"else",
		"while",
		"return",
	}
}
