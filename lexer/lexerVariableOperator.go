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
func ParseVariableAssigningExpression(key Token, expression Token, value Token, lexer *Lexer) (output any, err error) {
	exp := expression
	//if value.Type != NUMBER && value.Type != IDENTIFIER && value.Type != BOOL {
	//	if exp.Type != ASSIGN {
	//		return 0, fmt.Errorf("expected an integer, but got '%s'", value.Value)
	//	}
	//
	//}
	// if exp.Type != ASSIGN && value.Type != NUMBER && value.Type != STRING {
	// 	return 0, fmt.Errorf("non integers only support '=' operator for assigning")
	// }
	if key.Type != IDENTIFIER {
		return 0, fmt.Errorf("expected an identifier, but got '%s'", key.Value)
	}

	k, ok := std.Variables[key.Value]
	//if value.Type == BOOL && expression.Type == ASSIGN {
	//	std.Variables[key.Value] = value.Value == "true"
	//	return 0, err
	//}
	if !ok {
		return 0, fmt.Errorf("'%s' is not defined", key.Value)
	}

	//if reflect.TypeOf(k).String() != "int" {
	//	return 0, fmt.Errorf("can not do math operations on a non integer '%s'", key.Value)
	//}
	out, _, errs := MathExpressionTokensToEnd(value, lexer)
	if errs != nil {
		return 0, errs
	}
	o, errss := ParseArithmeticExpressions(out, lexer)
	if errss != nil {
		return 0, errss
	}
	//add mismatched type error message
	if reflect.TypeOf(o).String() != reflect.TypeOf(k.Value).String() {
		return 0, fmt.Errorf("mismatched types")
	}
	//vals, err := ParseExpression(value, lexer)
	if str, ok1 := o.(string); ok1 {
		if ConvertToTokenType(reflect.TypeOf(k.Value).String()) != STRING {
			return 0, fmt.Errorf("can not assign a non string to an string")
		}
		switch exp.Type {
		case ASSIGN:
			k.SetValue(str)
			std.Variables[key.Value] = k
			return str, nil
		case PLUS_ASSIGN:
			k.SetValue(k.Value.(string) + str)
			std.Variables[key.Value] = k
			return str, nil
		default:
			return "", fmt.Errorf("invalid operator for string assigning")
		}
	}

	if bol, ok1 := o.(bool); ok1 {
		if reflect.TypeOf(k.Value).String() == "int" {
			return 0, fmt.Errorf("can not assign a non boolean to an boolean")
		}
		if exp.Type != ASSIGN {
			return 0, fmt.Errorf("non integers only support '=' operator for assigning")
		}
		k.SetValue(bol)
		std.Variables[key.Value] = k
		return bol, nil
	}
	valI, ok := o.(int)

	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, nil
	}

	switch exp.Type {
	case PLUS_ASSIGN:
		k.SetValue(k.Value.(int) + valI)
		std.Variables[key.Value] = k
		return k.Value.(int), nil
	case SUBTRACT_ASSIGN:
		k.SetValue(k.Value.(int) - valI)
		std.Variables[key.Value] = k
		return k.Value.(int), nil
	case MULTIPLY_ASSIGN:
		k.SetValue(k.Value.(int) * valI)
		std.Variables[key.Value] = k
		return k.Value.(int), nil
	case DIVIDE_ASSIGN:
		k.SetValue(k.Value.(int) / valI)
		std.Variables[key.Value] = k
		return k.Value.(int), nil
	case MODULO_ASSIGN:
		k.SetValue(k.Value.(int) % valI)
		std.Variables[key.Value] = k
		return k.Value.(int), nil
	case ASSIGN:
		k.SetValue(valI)
		std.Variables[key.Value] = k
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
