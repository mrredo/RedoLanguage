package lexer

import (
	"RedoLanguage/std"
	"io"
	"log"
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
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	scanner scanner.Scanner
}

func NewLexer(input string) *Lexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanStrings |
		scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments
	return &Lexer{scanner: s}
}

func (l *Lexer) NextToken() Token {
	tok := l.scanner.Scan()
	for tok == scanner.Comment || tok == scanner.EOF {
		if tok == scanner.EOF {
			return Token{Type: EOF, Value: ""}
		}
		tok = l.scanner.Scan()
	}
	val := l.scanner.TokenText()

	switch tok {
	case scanner.Ident:
		if val == "var" {
			return Token{Type: VAR, Value: "var"}
		} else if val == "true" || val == "false" {
			return Token{Type: BOOL, Value: val}
		}
		return Token{Type: IDENTIFIER, Value: val}
	case '=':

		if l.scanner.Peek() == '=' {
			return Token{Type: EQUAL, Value: val}
		}

		return Token{Type: ASSIGN, Value: val}

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
func TestLexer() {
	lexer := NewLexer(`
println(add(1, 6))

`)
	for {
		// Parse the next function call
		funcName, args, err := ParseFunctionCall(lexer)

		if err == io.EOF {
			break // End of input
		}
		if err != nil {
			log.Fatal(err)
		}
		if funcName == "" {
			break
		}
		// Look up the function and call it with the arguments
		fn, ok := std.Functions[funcName]
		if !ok {

			log.Fatalf("undefined function '%s'", funcName)
		}
		fn(args...)
	}
}
