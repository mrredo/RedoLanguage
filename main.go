package main

import (
	"RedoLanguagev2/lexer"
	"fmt"
)

func main() {
	//cli.Execute()
	content := `
var hello = 10+10
print(hello+10)
`
	
	tokens := lexer.Tokenize(content)
	for _, v := range tokens {
		fmt.Printf("Type: %d, Value: %s\n", v.Type, v.Value)
	}
}
