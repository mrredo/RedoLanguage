package main

import (
	"RedoLanguagev2/cli"
	"RedoLanguagev2/lexer"
	"fmt"
)

func main() {
	cli.Execute()
	content := `
	10+10*(20*20)
10+10+10+"eeee"
`
	tokens := lexer.Tokenize(content)
	for _, v := range tokens {
		fmt.Printf("Type: %d, Value: %s\n", v.Type, v.Value)
	}
}
