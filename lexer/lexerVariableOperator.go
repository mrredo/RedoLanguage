package lexer

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
PSEUDO CODE:
if IsVariableExpression then
if IsVariableAssigningExpression() // += , -=
	expression := lexer.NextToken()
else
ParseOperator() // *, /, +, -

*/

/*
key = value
key += value
key -= value
key *= value
key /= value
key %= value
*/
func ParseVariableExpression(curT Token, lexer *Lexer) (key string, value any, err error) {
	operator := lexer.NextToken()
	//value := lexer.NextToken()
	switch operator.Type {
	case PLUS_ASSIGN:

	}
	return "", nil, err
}

/*
Only checks for +, -, *, /, %, =

it detects its a expression, then u can to lexer.NextToken()
*/
func IsVariableExpression(curT Token, lexer *Lexer) bool {
	pk := lexer.Scanner.Peek()

	return curT.Type == IDENTIFIER && (pk == '+' || pk == '-' || pk == '*' || pk == '/' || pk == '%' || pk == '=')
}
func ParseVariableAssigningExpression(key Token, expression Token, value Token, lexer *Lexer) (output int, err error) {
	exp := expression
	if value.Type != NUMBER {
		return 0, fmt.Errorf("expected an integer, but got '%s'", value.Value)
	}
	if key.Type != IDENTIFIER {
		return 0, fmt.Errorf("expected an identifier, but got '%s'", key.Value)
	}
	k, ok := Variables[key.Value]
	if !ok {
		return 0, fmt.Errorf("'%s' is not defined", key.Value)
	}
	if reflect.TypeOf(k).String() != "int" {
		return 0, fmt.Errorf("can not do math operations on a non integer '%s'", key.Value)
	}
	valI, err := strconv.Atoi(value.Value)

	if err != nil {

		return 0, err
	}
	switch exp.Type {
	case PLUS_ASSIGN:
		Variables[key.Value] = k.(int) + valI
		return k.(int) + valI, nil
	case SUBTRACT_ASSIGN:
		Variables[key.Value] = k.(int) - valI
		return k.(int) - valI, nil
	case MULTIPLY_ASSIGN:
		Variables[key.Value] = k.(int) * valI
		return k.(int) * valI, nil
	case DIVIDE_ASSIGN:
		Variables[key.Value] = k.(int) / valI
		return k.(int) / valI, nil
	case MODULO_ASSIGN:
		Variables[key.Value] = k.(int) % valI
		return k.(int) % valI, nil
	case ASSIGN:
		Variables[key.Value] = valI
		return valI, nil
	}
	return 0, nil
}
func ParseExpressionType(exp Token) TokenType {
	switch exp.Value {
	case "*=":
		return MULTIPLY_ASSIGN
	case "/=":
		return DIVIDE_ASSIGN
	case "+=":
		return PLUS_ASSIGN
	case "-=":
		return SUBTRACT_ASSIGN
	case "%=":
		return MODULO_ASSIGN
	case "=":
		return ASSIGN

	default:
		return 0
	}
}
