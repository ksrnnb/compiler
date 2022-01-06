package tokenizer

import "errors"

func keyword(token string) (KeywordType, error) {
	switch token {
	case "class":
		return Class, nil
	case "method":
		return Method, nil
	case "function":
		return Function, nil
	case "constructor":
		return Constructor, nil
	case "field":
		return Field, nil
	case "static":
		return Static, nil
	case "var":
		return Var, nil
	case "int":
		return Int, nil
	case "char":
		return Char, nil
	case "boolean":
		return Boolean, nil
	case "void":
		return Void, nil
	case "true":
		return True, nil
	case "false":
		return False, nil
	case "null":
		return Null, nil
	case "this":
		return This, nil
	case "let":
		return Let, nil
	case "do":
		return Do, nil
	case "if":
		return If, nil
	case "else":
		return Else, nil
	case "while":
		return While, nil
	case "return":
		return Return, nil
	}

	return 0, errors.New("keyword(): keyword is invalid")
}
