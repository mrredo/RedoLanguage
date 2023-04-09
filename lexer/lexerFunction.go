package lexer

import (
	"RedoLanguage/std"
	"errors"
	"fmt"
	"strconv"
)

func ParseFunctionCall(curT Token, sec Token, lexer *Lexer) (string, []interface{}, error) {
	var args []interface{}
	tok := curT

	if tok.Type == 0 {
		return "", nil, nil
	}
	if tok.Type != IDENTIFIER {
		return "", nil, errors.New("expected function name")
	}
	funcName := tok.Value
	if _, ok := std.Functions[funcName]; !ok {
		return funcName, nil, fmt.Errorf("undefined function '%s'", funcName)
	}
	tok = sec
	if tok.Type != LPAREN {
		return "", nil, errors.New("expected '(' after function name")
	}
	for {
		tok = lexer.NextToken()
		if tok.Type == RPAREN {
			break
		}
		if tok.Type == STRING {
			args = append(args, tok.Value[1:len(tok.Value)-1])

			tok = lexer.NextToken()
		}
		if tok.Type == NUMBER || tok.Type == IDENTIFIER || tok.Type == BOOL || tok.Type == LPAREN {
			out, l, errs := MathExpressionTokensToEndFunctionArgument(tok, lexer)
			if errs != nil {
				return "", nil, errs
			}
			o, errss := ParseArithmeticExpressions(out)
			if errss != nil {
				return "", nil, errss
			}
			tok = l
			//arg, err := ParseExpression(tok, lexer)
			//if err != nil {
			//	return "", nil, err
			//}

			args = append(args, o)
		}

		//tok = lexer.NextToken()
		if tok.Type == RPAREN {
			break
		} else if tok.Type != COMMA {
			return "", nil, errors.New("expected ',' or ')' after argument")
		}
	}
	//fmt.Println(args)
	return funcName, args, nil
}

func ParseExpression(tok Token, lexer *Lexer) (interface{}, error) {
	switch tok.Type {
	case 0:
		return nil, nil
	case BOOL:
		return tok.Value == "true", nil
	case NUMBER:
		return strconv.Atoi(tok.Value)
	case STRING:
		return tok.Value[1 : len(tok.Value)-1], nil
	case IDENTIFIER:
		if p := lexer.Scanner.Peek(); p != '(' {
			if val, ok := std.Variables[tok.Value]; ok {
				return val, nil
			}
			return nil, errors.New(fmt.Sprintf("'%s' is not defined", tok.Value))
		}

		//if tok.Value == "true" {
		//	return true, nil
		//} else if tok.Value == "false" {
		//	return false, nil
		//} else {

		// function call
		funcName := tok.Value
		tok := lexer.NextToken()
		if tok.Type != LPAREN {
			return nil, errors.New("expected '(' after function name")
		}
		args := make([]interface{}, 0)
		for {
			tok := lexer.NextToken()
			if tok.Type == RPAREN {
				break
			}
			arg, err := ParseExpression(tok, lexer)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
			tok = lexer.NextToken()
			if tok.Type == RPAREN {
				break
			} else if tok.Type != COMMA {
				return nil, errors.New("expected ',' or ')' after argument")
			}
		}
		fn, ok := std.Functions[funcName]
		if !ok {
			return nil, fmt.Errorf("undefined function '%s'", funcName)
		}
		return fn(args...), nil
		//}
	default:
		return nil, fmt.Errorf("unexpected token '%s'", tok.Value)
	}
}
func IsFunction(token Token, secondT Token, lexer *Lexer) bool {
	return token.Type == IDENTIFIER && secondT.Type == LPAREN
}

//func TestIsFunction() {
//	lx := NewLexer(`
//print()
//`)
//	c := lx.NextToken()
//	fmt.Println(IsFunction(c, lx))
//}
