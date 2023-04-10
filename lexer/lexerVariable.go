package lexer

import (
	"RedoLanguage/std"
	"errors"
	"fmt"
	"strconv"
)

func RuneToStr(ch rune) string {
	return fmt.Sprintf(`%q`, string(ch))
}
func ParseVariable(curToken Token, sec Token, lexer *Lexer) (key string, value interface{}, err error) {
	tok := curToken //
	keyT := sec
	//fmt.Println(tok)
	if tok.Type == 0 {
		return "", nil, nil
	}
	if tok.Type != VAR {
		return "", nil, errors.New("var err unknown, (if you got this error message you messed up real bad)")
	}
	if keyT.Type != IDENTIFIER {

		return "", nil, fmt.Errorf("expected an identifier, but got '%s'", keyT.Value)
	}
	if VariableExists(keyT.Value) {
		return "", nil, fmt.Errorf("'%s' is already declared", keyT.Value)
	}
	Eq := lexer.NextToken()
	//fmt.Println(Eq)
	if Eq.Type != ASSIGN {
		return "", nil, fmt.Errorf("'=' sign is expected after the '%s'", keyT.Value)
	}
	valT := lexer.NextToken()
	//
	//if valT.Type == STRING {
	//	parsedVal, err := ParseExpression(valT, lexer)
	//	if err != nil {
	//		return "", nil, err
	//	}
	//
	//	std.Variables[keyT.Value] = parsedVal
	//	return keyT.Value, parsedVal, nil
	//}

	//parsedVal, err := ParseExpression(valT, lexer)
	if err != nil {
		return "", nil, err
	}
	out, _, errs := MathExpressionTokensToEnd(valT, lexer)
	if errs != nil {
		return "", nil, errs
	}
	o, errss := ParseArithmeticExpressions(out)
	if errss != nil {
		return "", 0, errss
	}
	//bol, ok1 := o.(bool)
	//if ok1 {
	//	std.Variables[keyT.Value] = bol
	//	return keyT.Value, bol, nil
	//}
	//valI, ok := o.(int)
	//if !ok {
	//	return "", nil, fmt.Errorf("error parsing value to int")
	//}
	std.Variables[keyT.Value] = o
	return keyT.Value, o, nil

}
func VariableExists(name string) bool {
	return std.Variables[name] != nil
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
	if val, ok := std.Variables[token.Value]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("undefined identifier: %s", token.Value)
}
