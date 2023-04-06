package main

import (
	"RedoLanguage/interpreter"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		inp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if inp == "exit" {
			break
		}
		inp = strings.TrimRight(inp, "\n")

		interpreter.Interpret(inp)
		fmt.Println()
		fmt.Println()

	}

}
