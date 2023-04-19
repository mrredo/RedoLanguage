package lexer

func IsIfStatement(c, s Token, lx *Lexer) (bool, *If, error) {
	switch c.Type {
	case IF, ELSE, ELSE_IF:

		break

	default:
		return false, nil, nil
	}
	lx.CurrentNestingLevel++
	lx.CurrentPosition ++
	lx.IfPositions[lx.CurrentNestingLevel] = If{
		Condition: "",
		Position: lx.CurrentPosition,
		NestingLevel: lx.CurrentNestingLevel,
		ElseIfs: []elseIf{},
	}
	return false, nil, nil
	
}