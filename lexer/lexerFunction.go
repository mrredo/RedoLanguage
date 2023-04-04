package lexer

import (
	"RedoLanguage/std"
	"errors"
	"fmt"
	"strconv"
)

func ParseFunctionCall(lexer *Lexer) (string, []interface{}, error) {
	var args []interface{}
	tok := lexer.NextToken()

	if tok.Type == 0 {
		return "", nil, nil
	}
	if tok.Type != IDENTIFIER {
		return "", nil, errors.New("expected function name")
	}
	funcName := tok.Value
	tok = lexer.NextToken()
	if tok.Type != LPAREN {
		return "", nil, errors.New("expected '(' after function name")
	}
	for {
		tok = lexer.NextToken()
		if tok.Type == RPAREN {
			break
		}
		arg, err := ParseExpression(tok, lexer)
		if err != nil {
			return "", nil, err
		}
		args = append(args, arg)
		tok = lexer.NextToken()
		if tok.Type == RPAREN {
			break
		} else if tok.Type != COMMA {
			return "", nil, errors.New("expected ',' or ')' after argument")
		}
	}
	return funcName, args, nil
}

func ParseExpression(tok Token, lexer *Lexer) (interface{}, error) {
	switch tok.Type {
	case 0:
		return nil, nil
	case NUMBER:
		return strconv.Atoi(tok.Value)
	case STRING:
		return tok.Value[1 : len(tok.Value)-1], nil
	case IDENTIFIER:
		if tok.Value == "true" {
			return true, nil
		} else if tok.Value == "false" {
			return false, nil
		} else {
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
		}
	default:
		return nil, fmt.Errorf("unexpected token '%s'", tok.Value)
	}
}
func IsFunction(token Token, lexer *Lexer) bool {
	par := lexer.scanner.Peek()
	if token.Type == IDENTIFIER && par == '(' {
		return true
	}
	return false
}
func TestIsFunction() {
	lx := NewLexer(`
print()
`)
	c := lx.NextToken()
	fmt.Println(IsFunction(c, lx))
}
