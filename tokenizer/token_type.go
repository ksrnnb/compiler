package tokenizer

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/ksrnnb/compiler/utils"
)

func tokenType(token string) (TokenType, error) {
	if isKeyword(token) {
		return Keyword, nil
	}

	if isSymbol(token) {
		return Symbol, nil
	}

	if isIntConst(token) {
		return IntConst, nil
	}

	if isStringConst(token) {
		return StringConst, nil
	}

	if isIdentifier(token) {
		return Identifier, nil
	}

	return 0, errors.New("tokenType: token type is invalid")
}

func isKeyword(token string) bool {
	return utils.IsInSlice(getKeywordTokens(), token)
}

func isSymbol(token string) bool {
	return utils.IsInSlice(getKeywordTokens(), token)
}

func isIntConst(token string) bool {
	// 整数であればIntConst => err == nil
	intConstValue, err := strconv.Atoi(token)

	if err != nil {
		return false
	}

	// 32767 = 2^15 - 1 => 16bitで表す整数の上限値
	if intConstValue < 0 || intConstValue > 32767 {
		return false
	}

	return true
}

func isStringConst(token string) bool {
	// string => ダブルクォーテンションで囲まれる。
	r := regexp.MustCompile(`^"[^"\n].*"$`)
	return r.MatchString(token)
}

func isIdentifier(token string) bool {
	r := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	return r.MatchString(token)
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

func getSymbolTokens() []string {
	return []string{
		"{",
		"}",
		"(",
		")",
		"[",
		"]",
		".",
		",",
		";",
		"+",
		"-",
		"*",
		"/",
		"&",
		"|",
		"<",
		">",
		"=",
		"~",
	}
}
