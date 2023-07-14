package lexer

import (
	"RedoLanguagev2/types"
	"go/scanner"
	gtoken "go/token"
	"strings"
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
	fset := gtoken.NewFileSet()
	file := fset.AddFile("", -1, len(line))
	s.Init(file, []byte(line), nil, 0)

	var currentToken string
	inQuotes := false

	for {
		/*pos,*/ _, tok, lit := s.Scan()
		tokenText := lit

		if tok == gtoken.EOF {
			break
		}

		if tok == gtoken.STRING {
			if inQuotes {
				// Add the complete string token
				token := types.Token{
					Type:  types.String,
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

func getTokenType(tok gtoken.Token) types.TokenType {
	switch tok {
	case gtoken.IDENT:
		switch tok.String() {
		case "var":
			return types.Var
		case "if":
			return types.If

		}
		return types.Identifier
	case gtoken.ADD_ASSIGN:
		return types.PlusAssign
	case gtoken.SUB_ASSIGN:
		return types.MinusAssign
	case gtoken.MUL_ASSIGN:
		return types.MultAssign
	case gtoken.QUO_ASSIGN:
		return types.PlusAssign
	case gtoken.STRING:
		return types.String
	case gtoken.INT, gtoken.FLOAT:
		return types.Number
	case gtoken.ADD:
		return types.Plus
	case gtoken.SUB:
		return types.Minus
	case gtoken.MUL:
		return types.Multiply
	case gtoken.QUO:
		return types.Divide
	case gtoken.LPAREN:
		return types.LeftParenthesis
	case gtoken.RPAREN:
		return types.RightParenthesis
	case gtoken.LBRACK:
		return types.LBrack
	case gtoken.SEMICOLON:
		return types.SemiColon
	case gtoken.RBRACK:
		return types.RBrack
	default:
		if isNumberToken(tok) {
			return types.Number
		}
		return types.Unknown
	}
}

func isNumberToken(tok gtoken.Token) bool {
	return tok == gtoken.INT || tok == gtoken.FLOAT
}
