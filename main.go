package main

import (
	"RedoLanguage/interpreter"
	"RedoLanguage/reader"
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

	//fmt.Println(err.NewTypeError(scanner.Position{
	//	Filename: "main.rd",
	//	Offset:   100,
	//	Line:     10,
	//	Column:   10,
	//}))
	fmt.Println(reader.ReadFileContent("hello.rd"))
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
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
