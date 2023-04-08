package main

import (
	"RedoLanguage/interpreter"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//lx.TestMath()
	//n := lx.NewLexer(`(10+10)*10`)
	//l := lx.NewLexer("10 + 10 * 10")
	//
	//fmt.Println(lx.IsMathExpression(l.NextToken(), l.NextToken(), l))
	//os.Exit(1)
	//fmt.Print(lx.ParseArithmeticExpressions(n.NextToken(), n.NextToken(), n))
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
