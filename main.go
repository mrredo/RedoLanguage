package main

import (
	"RedoLanguagev2/ast"
	"RedoLanguagev2/lexer"
	"encoding/json"
	"fmt"
)

func main() {
	//cli.Execute()
	content := `

🌍
var hello =10
`

	tokens := lexer.Tokenize(content)
	for _, v := range tokens {
		fmt.Printf("Type: %d, Value: %s\n", v.Type, v.Value)
	}
	astStr := ast.ParseTokens(tokens)
	jsonData, err := json.Marshal(astStr)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Convert JSON byte array to string
	jsonString := string(jsonData)

	// Print the JSON string
	fmt.Println(jsonString)
}
