package lexer

import "fmt"

func PrintTokens(lexer *Lexer) {
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break // Stop looping when the end of input is reached
		}
		fmt.Printf("%v: %v\n", token.Type, token.Value)
	}
}
