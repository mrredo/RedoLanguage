package main

import (
	"RedoLanguage/interpreter"
	"RedoLanguage/reader"
	"bufio"
	"fmt"
	"log"
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

	//fmt.Println(err.NewTypeError(scanner.Position{
	//	Filename: "main.rd",
	//	Offset:   100,
	//	Line:     10,
	//	Column:   10,
	//}))

	if len(os.Args) <= 1 {

		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">>> ")
			inp, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			inp = strings.TrimRight(inp, "\n")

			interpreter.Interpret(inp, "interpret.rd")
			fmt.Println()
			fmt.Println()

		}
	}
	fileName := os.Args[1]
	str, err := reader.ReadFileContent(fileName)
	if err != nil {
		log.Println(err.Error())
		return
	}
	interpreter.Interpret(str, fileName)

}
