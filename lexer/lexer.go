package lexer

import (
	"RedoLanguage/err"
	"regexp"
	"strings"
	"text/scanner"
)

type TokenType int

const (
	EOF                   TokenType = iota // end of file
	LPAREN                                 // (
	RPAREN                                 // )
	COMMA                                  // ,
	IDENTIFIER                             // variable identifier
	NUMBER                                 // numeric literal
	STRING                                 // string literal
	EQUAL                                  // ==
	ASSIGN                                 // =
	VAR                                    // var keyword
	BOOL                                   // boolean literal
	MODULO                                 // %
	DIVIDE                                 // /
	PLUS                                   // +
	MULTIPLY                               // *
	SUBTRACT                               // -
	PLUS_ASSIGN                            // +=
	SUBTRACT_ASSIGN                        // -=
	MULTIPLY_ASSIGN                        // *=
	DIVIDE_ASSIGN                          // /=
	MODULO_ASSIGN                          // %=
	LEFT_SHIFT_ASSIGN                      // <<=
	RIGHT_SHIFT_ASSIGN                     // >>=
	BITWISE_AND_ASSIGN                     // &=
	BITWISE_XOR_ASSIGN                     // ^=
	PLUS_PLUS                              // ++
	SUBTRACT_SUBTRACT                      // --
	BITWISE_XOR                            // ^
	LEFT_SHIFT                             // <<
	RIGHT_SHIFT                            // >>
	BITWISE_AND                            // &
	EQUAL_TO                               // ==
	NOT_EQUAL_TO                           // !=
	LESS_THAN                              // <
	LESS_THAN_OR_EQUAL                     // <=
	GREATER_THAN                           // >
	GREATER_THAN_TO_EQUAL                  // >=
	AND                                    // &&
	SEMICOLON                              // ;
	NEW_LINE                               // \n
	BITWISE_OR                             // |
	OR                                     // ||
	ILLEGAL

	SEMICOLON_ERROR // does not have a string but is
)

var OperatorNumToString = map[TokenType]string{
	LPAREN:                "(",
	RPAREN:                ")",
	EQUAL:                 "==",
	MODULO:                "%",
	DIVIDE:                "/",
	PLUS:                  "+",
	MULTIPLY:              "*",
	SUBTRACT:              "-",
	PLUS_PLUS:             "++",
	SUBTRACT_SUBTRACT:     "--",
	BITWISE_XOR:           "^",
	LEFT_SHIFT:            "<<",
	RIGHT_SHIFT:           ">>",
	BITWISE_AND:           "&",
	EQUAL_TO:              "==",
	NOT_EQUAL_TO:          "!=",
	LESS_THAN:             "<",
	LESS_THAN_OR_EQUAL:    "<=",
	GREATER_THAN:          ">",
	GREATER_THAN_TO_EQUAL: ">=",
	AND:                   "&&",
	BITWISE_OR:            "|",
	OR:                    "||",
}

var numReg = regexp.MustCompile(`\d`)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	Scanner       scanner.Scanner
	Input         string
	semiColonLine int
	SemErr        error
}

func NewLexer(input string) *Lexer {
	var s scanner.Scanner

	s.Init(strings.NewReader(input))
	s.Filename = "interpreter.rd"
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanStrings |
		scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments

	return &Lexer{Scanner: s, Input: input, semiColonLine: 0}
}

func (l *Lexer) NextToken() Token {
	tok := l.Scanner.Scan()
	//var lastLine int
	for tok == scanner.Comment || tok == scanner.EOF {
		//if l.Scanner.Pos().Line != lastLine { // check if the line number has changed
		//	lastLine = l.Scanner.Pos().Line // update lastLine to the current line number
		//}
		if tok == scanner.EOF {
			return Token{Type: EOF, Value: ""}
		}
		tok = l.Scanner.Scan()
	}
	if tok == ';' && l.semiColonLine != l.Scanner.Pos().Line {
		l.semiColonLine = l.Scanner.Pos().Line + 1
		return Token{Type: SEMICOLON, Value: ";"}
	} else if l.semiColonLine != l.Scanner.Pos().Line {
		l.SemErr = err.NewSemiColonError(l.Scanner.Pos())
		return Token{Type: ILLEGAL, Value: ";"}
	}
	val := l.Scanner.TokenText()
	// if l.Scanner.Pos().Line != l.curLine {
	// 	l.curLine = l.Scanner.Pos().Line
	// 	return Token{Type: NEW_LINE, Value: "\n"}

	// }
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
	case '\n': // \n
		return Token{Type: NEW_LINE, Value: "\n"}
	case '&':

		if p := l.Scanner.Peek(); p == '&' {
			l.NextToken()
			return Token{Type: OR, Value: "&&"}
		}
		return Token{Type: BITWISE_AND, Value: "&"}

	case '|':
		{
			if p := l.Scanner.Peek(); p == '|' {
				l.NextToken()
				return Token{Type: AND, Value: "||"}
			}
			return Token{Type: BITWISE_OR, Value: "|"}
		}
	case ';':
		return Token{Type: SEMICOLON, Value: ";"}
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
		} else if numReg.MatchString(string(l.Scanner.Peek())) {
			return Token{Type: NUMBER, Value: "-" + l.NextToken().Value}
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

		return Token{Type: ILLEGAL, Value: val}
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
