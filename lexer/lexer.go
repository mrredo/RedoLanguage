package lexer

import (
	"strings"
	"text/scanner"
)

type TokenType int

const (
	EOF TokenType = iota
	LPAREN
	RPAREN
	COMMA
	IDENTIFIER
	NUMBER
	STRING
	EQUAL
	ASSIGN
	VAR
	BOOL
	MODULO
	DIVIDE
	PLUS
	MULTIPLY
	SUBTRACT
	PLUS_ASSIGN
	SUBTRACT_ASSIGN
	MULTIPLY_ASSIGN
	DIVIDE_ASSIGN
	MODULO_ASSIGN
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	Scanner scanner.Scanner
}

func NewLexer(input string) *Lexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanStrings |
		scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments
	return &Lexer{Scanner: s}
}

func (l *Lexer) NextToken() Token {
	tok := l.Scanner.Scan()
	for tok == scanner.Comment || tok == scanner.EOF {
		if tok == scanner.EOF {
			return Token{Type: EOF, Value: ""}
		}
		tok = l.Scanner.Scan()
	}
	val := l.Scanner.TokenText()

	switch tok {
	case scanner.Ident:
		if val == "var" {
			return Token{Type: VAR, Value: "var"}
		} else if val == "true" || val == "false" {
			return Token{Type: BOOL, Value: val}
		}
		return Token{Type: IDENTIFIER, Value: val}
	case '=':

		if l.Scanner.Peek() == '=' {
			l.NextToken()
			return Token{Type: EQUAL, Value: "=="}
		}

		return Token{Type: ASSIGN, Value: val}
	case '+':
		if l.Scanner.Peek() == '=' {
			l.NextToken()
			return Token{Type: PLUS_ASSIGN, Value: "+="}
		}
		return Token{Type: PLUS, Value: val}
	case '-':
		if l.Scanner.Peek() == '=' {
			l.NextToken()
			return Token{Type: SUBTRACT_ASSIGN, Value: "-="}
		}
		return Token{Type: SUBTRACT, Value: val}
	case '*':
		if l.Scanner.Peek() == '=' {
			l.NextToken()
			return Token{Type: MULTIPLY_ASSIGN, Value: "*="}
		}
		return Token{Type: MULTIPLY, Value: val}

	case '/':
		if l.Scanner.Peek() == '=' {
			l.NextToken()
			return Token{Type: DIVIDE_ASSIGN, Value: "/="}
		}
		return Token{Type: DIVIDE, Value: val}
	case '%':
		if l.Scanner.Peek() == '=' {
			l.NextToken()
			return Token{Type: MODULO_ASSIGN, Value: "%="}
		}
		return Token{Type: MODULO, Value: val}
	case ',':
		return Token{Type: COMMA, Value: val}
	case '(':
		return Token{Type: LPAREN, Value: val}
	case ')':
		return Token{Type: RPAREN, Value: val}
	case scanner.Float:
		return Token{Type: NUMBER, Value: val}
	case scanner.Int:
		return Token{Type: NUMBER, Value: val}
	case scanner.String:
		return Token{Type: STRING, Value: val}

	default:

		return Token{Type: IDENTIFIER, Value: val}
	}
}

//func (l *Lexer) PeekToken() Token {
//	currentPos := l.Scanner.Pos()
//
//	tok := l.NextToken()
//	l.Scanner.Peek()
//	return tok
//}

func TestLexer() {
	//	lexer := NewLexer(`
	//println(add(1, 6))
	//
	//`)
	//for {
	//	// Parse the next function call
	//	funcName, args, err := ParseFunctionCall(lexer)
	//
	//	if err == io.EOF {
	//		break // End of input
	//	}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	if funcName == "" {
	//		break
	//	}
	//	// Look up the function and call it with the arguments
	//	fn, ok := std.Functions[funcName]
	//	if !ok {
	//
	//		log.Fatalf("undefined function '%s'", funcName)
	//	}
	//	fn(args...)
	//}
}
