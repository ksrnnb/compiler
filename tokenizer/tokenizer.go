package tokenizer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

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
	scanner              *bufio.Scanner
	isDone               bool
	currentToken         string
	currentTokens        []string
	isInMultiLineComment bool
}

func NewTokenizer(input io.Reader) *Tokenizer {
	sc := bufio.NewScanner(input)
	tokenizer := &Tokenizer{scanner: sc, isDone: false}

	tokenizer.output()
	return tokenizer
}

func (t *Tokenizer) output() {
	// TODO: ***T.xmlファイルの作成
	for {
		if !t.HasMoreTokens() {
			break
		}

		t.Advance()
		fmt.Println(t.currentToken)
	}
}

func (t Tokenizer) HasMoreTokens() bool {
	return !t.isDone
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

		trimmedLine := utils.TrimSpaceAndTab(line)

		if len(trimmedLine) == 0 {
			t.Advance()
			return
		}

		t.currentTokens = strings.Split(trimmedLine, " ")
		t.Advance()
		return
	}

	t.currentToken = t.currentTokens[0]

	// 複数行コメント
	if t.isInMultiLineComment {
		t.skipMultiLineComment()
		t.Advance()
		return
	}

	// 複数行コメント
	if strings.HasPrefix(t.currentToken, "/*") {
		t.isInMultiLineComment = true
		t.Advance()
		return
	}

	// コメントの場合は、残りを無視して次の行へ
	if strings.HasPrefix(t.currentToken, "//") {
		t.currentTokens = []string{}
		t.Advance()
		return
	}

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

func (t *Tokenizer) skipMultiLineComment() {
	for i, token := range t.currentTokens {
		index := strings.Index(token, "*/")
		if index == -1 {
			continue
		}

		t.isInMultiLineComment = false
		afterCommentIndex := index + len("*/")

		// コメントのあとにスペースがあって、その後に式が続く場合
		if len(t.currentTokens) > i {
			t.currentTokens = t.currentTokens[i+1:]
		} else {
			t.currentTokens = []string{}
		}

		// コメントの後にもスペースなしで式が続く場合
		if len(token) > afterCommentIndex {
			newToken := utils.TrimSpaceAndTab(token[afterCommentIndex:])
			t.currentTokens = append([]string{newToken}, t.currentTokens...)
			return
		}

		return
	}

	// " */ " がみつからない場合は次の行を読み込むため、リセットする
	t.currentTokens = []string{}
}
