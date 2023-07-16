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
		/*pos,*/ pos, tok, lit := s.Scan()

		tokenText := lit

		if tok == gtoken.EOF {
			break
		}
		posi := fset.Position(pos)

		if tok == gtoken.STRING {
			if inQuotes {
				// Add the complete string token
				token := types.Token{
					Type:  types.String,
					Value: currentToken + `"`,
					Pos:   posi,
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
			ty, text := getTokenType(tok, tokenText)
			if text == "" {
				token := types.Token{
					Type:  ty,
					Value: tokenText,
					Pos:   posi,
				}
				tokens = append(tokens, token)
			} else {
				token := types.Token{
					Type:  ty,
					Value: text,
					Pos:   posi,
				}
				tokens = append(tokens, token)
			}

		}
	}

	return tokens
}

func getTokenType(tok gtoken.Token, tokText string) (types.TokenType, string) {
	ty := types.TokenType(0)
	text := ""
	switch tok {
	case gtoken.VAR:
		ty = types.Var
	case gtoken.IDENT /*, gtoken.ILLEGAL*/ :
		switch tokText {
		case "var":
			ty = types.Var

		case "if":
			ty = types.If

		}
		ty = types.Identifier
	case gtoken.ADD_ASSIGN:
		ty = types.PlusAssign
	case gtoken.SUB_ASSIGN:
		ty = types.MinusAssign
	case gtoken.MUL_ASSIGN:
		ty = types.MultAssign
	case gtoken.QUO_ASSIGN:
		ty = types.PlusAssign
	case gtoken.STRING:
		ty = types.String
	case gtoken.INT, gtoken.FLOAT:
		ty = types.Number
	case gtoken.ADD:
		text = "+"
		ty = types.Plus
	case gtoken.SUB:
		ty = types.Minus
	case gtoken.MUL:
		ty = types.Multiply
	case gtoken.QUO:
		ty = types.Divide
	case gtoken.LPAREN:
		ty = types.LeftParenthesis
	case gtoken.RPAREN:
		ty = types.RightParenthesis
	case gtoken.LBRACK:
		ty = types.LBrack
	case gtoken.SEMICOLON:
		ty = types.SemiColon
	case gtoken.RBRACK:
		ty = types.RBrack
	default:
		if isNumberToken(tok) {
			ty = types.Number
		}
		ty = types.Unknown
	}
	return ty, text
}

func isNumberToken(tok gtoken.Token) bool {
	return tok == gtoken.INT || tok == gtoken.FLOAT
}
