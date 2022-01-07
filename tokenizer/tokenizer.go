package tokenizer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
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
	scanner       *bufio.Scanner
	isDone        bool
	currentToken  string
	currentTokens []string
}

func NewTokenizer(input io.Reader) *Tokenizer {
	sc := bufio.NewScanner(input)
	return &Tokenizer{scanner: sc, isDone: false}
}

func (t Tokenizer) HasMoreTokens() bool {
	return t.isDone
}

func (t *Tokenizer) Advance() {
	if t.isDone {
		return
	}

	if len(t.currentTokens) == 0 {
		if !t.scanner.Scan() {
			t.isDone = true
			return
		}

		line := t.scanner.Text()
		trimmedLine := strings.Trim(line, " ")
		if len(trimmedLine) == 0 {
			t.Advance()
			return
		}

		t.currentTokens = strings.Split(trimmedLine, " ")
		t.Advance()
		return
	}

	token := t.currentTokens[0]

	// コメントの場合は、残りを無視して次の行へ
	if token == "//" {
		t.currentTokens = []string{}
		t.Advance()
		return
	}

	t.currentToken = token

	if len(t.currentTokens) == 1 {
		t.currentTokens = []string{}
		return
	}

	t.currentTokens = t.currentTokens[1:]
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

func (t Tokenizer) Identifier() (string, error) {
	tokenType, err := t.TokenType()
	if err != nil {
		return "", fmt.Errorf("Identifier(): token type should be identifier, %v", err)
	}

	if tokenType != Identifier {
		return "", errors.New("Identifier(): token type should be identifier")
	}

	return t.currentToken, nil
}

func (t Tokenizer) IntVal() (int, error) {
	tokenType, err := t.TokenType()
	if err != nil {
		return 0, fmt.Errorf("IntVal(): token type should be IntConst, %v", err)
	}

	if tokenType != IntConst {
		return 0, errors.New("IntVal(): token type should be IntConst")
	}

	intValue, err := strconv.Atoi(t.currentToken)

	if err != nil {
		return 0, fmt.Errorf("IntVal(): error while parsing int, %v", err)
	}
	return intValue, nil
}

func (t Tokenizer) StringVal() (string, error) {
	tokenType, err := t.TokenType()
	if err != nil {
		return "", fmt.Errorf("StringVal(): token type should be StringConst, %v", err)
	}

	if tokenType != StringConst {
		return "", errors.New("StringVal(): token type should be StringConst")
	}

	return t.currentToken, nil
}
