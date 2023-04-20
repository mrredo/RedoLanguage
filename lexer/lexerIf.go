package lexer

import "fmt"

func IsIfStatement(c, s Token, lx *Lexer) (bool, error) {
	switch c.Type {
	case IF, ELSE, ELSE_IF:
		break

	default:
		return false, nil
	}
	lx.CurrentNestingLevel++
	lx.CurrentPosition++
	var e *If = &If{
		Position:     lx.CurrentPosition,
		NestingLevel: lx.CurrentNestingLevel,
	}
	fmt.Println(e)
	return false, nil

}
