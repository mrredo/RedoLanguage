package lexer

import (
	"strings"
	"text/scanner"
)

type TokenType int

const (
	Number TokenType = iota
	Plus
	Minus
	Multiply
	Divide
	LeftParenthesis
	RightParenthesis
	NewLine
	STRING
	Unknown
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) []Token {
	var tokens []Token
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		lineTokens := tokenizeLine(line)
		tokens = append(tokens, lineTokens...)
		tokens = append(tokens, Token{Type: NewLine, Value: "\n"})
	}

	return tokens
}

func tokenizeLine(line string) []Token {
	var tokens []Token
	var s scanner.Scanner
	s.Init(strings.NewReader(line))
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.SkipComments

	var currentToken string
	inQuotes := false

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokenText := s.TokenText()

		if tok == '"' {
			if inQuotes {
				// Add the complete string token
				token := Token{
					Type:  STRING,
					Value: currentToken + `"`,
				}
				tokens = append(tokens, token)

				currentToken = ""
				inQuotes = false
				continue
			} else {
				inQuotes = true
			}
		}

		if inQuotes {
			currentToken += tokenText
		} else {
			token := Token{
				Type:  getTokenType(tok),
				Value: tokenText,
			}
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func getTokenType(tok rune) TokenType {
	switch tok {
	case scanner.String:
		return STRING
	case scanner.Int:
		return Number
	case '+':
		return Plus
	case '-':
		return Minus
	case '*':
		return Multiply
	case '/':
		return Divide
	case '(':
		return LeftParenthesis
	case ')':
		return RightParenthesis
	default:
		//if isNumber(tok) {
		//	return Number
		//}
		return Unknown
	}
}

func isNumber(tok rune) bool {
	return (tok >= '0' && tok <= '9') || tok == '.'
}
