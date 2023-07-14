package lexer

import (
	"RedoLanguagev2/types"
	"strings"
	"text/scanner"
)



func Tokenize(input string) []types.Token {
	var tokens []types.Token
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		lineTokens := tokenizeLine(line)
		tokens = append(tokens, lineTokens...)
		tokens = append(tokens, types.Token{Type: types.NewLine, Value: "\n"})
	}

	return tokens
}

func tokenizeLine(line string) []types.Token {
	var tokens []types.Token
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
				token := types.Token{
					Type:  types.STRING,
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
			token := types.Token{
				Type:  getTokenType(tok),
				Value: tokenText,
			}
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func getTokenType(tok rune) types.TokenType {
	switch tok {
	case scanner.Identifier: 
	
	case scanner.String:
		return types.STRING
	case scanner.Int:
		return types.Number
	case '+':
		return types.Plus
	case '-':
		return types.Minus
	case '*':
		return types.Multiply
	case '/':
		return types.Divide
	case '(':
		return types.LeftParenthesis
	case ')':
		return types.RightParenthesis
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
