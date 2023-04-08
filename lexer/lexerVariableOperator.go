package lexer

import (
	"RedoLanguage/std"
	"fmt"
	"reflect"
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

/*
Only checks for +, -, *, /, %, =

it detects its a expression, then u can to lexer.NextToken()
*/
func IsVariableExpression(curT Token, exp Token, lexer *Lexer) bool {
	//pk := lexer.Scanner.Peek()

	return curT.Type == IDENTIFIER && (exp.Type == ASSIGN || exp.Type == PLUS_ASSIGN || exp.Type == SUBTRACT_ASSIGN || exp.Type == MULTIPLY_ASSIGN || exp.Type == DIVIDE_ASSIGN || exp.Type == MODULO_ASSIGN) //pk != '(' //(pk == '+' || pk == '-' || pk == '*' || pk == '/' || pk == '%' || pk == '=')
}
func ParseVariableAssigningExpression(key Token, expression Token, value Token, lexer *Lexer) (output int, err error) {
	exp := expression
	if value.Type != NUMBER && value.Type != IDENTIFIER && value.Type != BOOL {
		if exp.Type != ASSIGN {
			return 0, fmt.Errorf("expected an integer, but got '%s'", value.Value)
		}

	}
	if exp.Type != ASSIGN && value.Type == BOOL {
		return 0, fmt.Errorf("booleans only support '=' operator for assigning")
	}
	if key.Type != IDENTIFIER {
		return 0, fmt.Errorf("expected an identifier, but got '%s'", key.Value)
	}

	k, ok := std.Variables[key.Value]
	if value.Type == BOOL && expression.Type == ASSIGN {
		std.Variables[key.Value] = value.Value == "true"
		return 0, err
	}
	if !ok {
		return 0, fmt.Errorf("'%s' is not defined", key.Value)
	}
	if reflect.TypeOf(k).String() != "int" {
		return 0, fmt.Errorf("can not do math operations on a non integer '%s'", key.Value)
	}
	vals, err := ParseExpression(value, lexer)

	valI, ok := vals.(int)

	if err != nil {

		return 0, err
	}
	if !ok {
		return 0, nil
	}
	switch exp.Type {
	case PLUS_ASSIGN:
		std.Variables[key.Value] = k.(int) + valI
		return k.(int) + valI, nil
	case SUBTRACT_ASSIGN:
		std.Variables[key.Value] = k.(int) - valI
		return k.(int) - valI, nil
	case MULTIPLY_ASSIGN:
		std.Variables[key.Value] = k.(int) * valI
		return k.(int) * valI, nil
	case DIVIDE_ASSIGN:
		std.Variables[key.Value] = k.(int) / valI
		return k.(int) / valI, nil
	case MODULO_ASSIGN:
		//fmt.Println(k.(int) % valI)
		std.Variables[key.Value] = k.(int) % valI
		return k.(int) % valI, nil
	case ASSIGN:
		std.Variables[key.Value] = valI
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
