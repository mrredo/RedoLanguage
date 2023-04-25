package lexer

import (
	"RedoLanguage/std"
	"errors"
	"fmt"
)

func RuneToStr(ch rune) string {
	return fmt.Sprintf(`%q`, string(ch))
}
func ParseVariable(curToken Token, sec Token, lexer *Lexer) (key string, value interface{}, err error) {
	tok := curToken //
	keyT := sec
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
	o, errss := ParseArithmeticExpressions(out, lexer)
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
	v := std.Variables[keyT.Value]
	v.SetValue(o)
	v.Key = keyT.Value
	v.NestingLevel = lexer.CurrentNestingLevel
	v.Position = lexer.CurrentPosition
	std.Variables[keyT.Value] = v
	return keyT.Value, o, nil

}
func VariableExists(name string) bool {
	return std.Variables[name].Key != ""
}
