package lexer

import "fmt"

func HandleIfStatement(c, s Token, lx *Lexer) (bool, error) {
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
func IsIfStatement(c Token) bool {
	return c.Type == IF || c.Type == ELSE || c.Type == ELSE_IF
}
