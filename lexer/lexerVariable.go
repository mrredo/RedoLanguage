package lexer

import (
	"errors"
	"fmt"
	"strconv"
)

var Variables = map[string]any{
	"hello": 10,
}

func RuneToStr(ch rune) string {
	return fmt.Sprintf(`%q`, string(ch))
}
func ParseVariable(curToken Token, lexer *Lexer) (key string, value interface{}, err error) {
	tok := curToken //
	keyT := lexer.NextToken()
	//fmt.Println(tok)
	if tok.Type == 0 {
		return "", nil, nil
	}
	if tok.Type != VAR {
		return "", nil, errors.New("var err unkown")
	}
	if keyT.Type != IDENTIFIER {
		return "", nil, errors.New(fmt.Sprintf("'%s' must be an identifier", keyT.Value))
	}
	Eq := lexer.NextToken()
	//fmt.Println(Eq)
	if Eq.Type != ASSIGN {
		return "", nil, errors.New(fmt.Sprintf("'=' sign is expected after the '%s'", keyT.Value))
	}
	valT := lexer.NextToken()
	parsedVal, err := ParseExpression(valT, lexer)
	if err != nil {
		return "", nil, err
	}
	if VariableExists(keyT.Value) {
		return "", nil, errors.New(fmt.Sprintf("'%s' is already declared", keyT.Value))
	}
	Variables[keyT.Value] = parsedVal
	return keyT.Value, parsedVal, nil

}
func VariableExists(name string) bool {
	return Variables[name] != nil
}
func parseVariableValue(token Token) (interface{}, error) {
	switch token.Type {
	case 0:
		return nil, nil
	case NUMBER:
		return strconv.Atoi(token.Value)
	case STRING:
		return token.Value[1 : len(token.Value)-1], nil
	case BOOL:
		if token.Value == "true" {
			return true, nil
		}
		return false, nil

	case IDENTIFIER:
		return parseIdentifier(token)
	default:
		/*
		   TODO: {
		   1. check if identifier exists and if does return identifier value
		   2. check if function then take function output if function returns nothing throw err

		   }




		*/
		return nil, nil

	}

}
func parseIdentifier(token Token) (interface{}, error) {
	if val, ok := Variables[token.Value]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("undefined identifier: %s", token.Value)
}
