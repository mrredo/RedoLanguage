package lexer

import (
	"RedoLanguage/err"
	"fmt"
)

type If struct {
	Condition    string
	Position     int
	NestingLevel int
}

func (ifs *If) Output() (bool, error) {
	lx := NewLexer(fmt.Sprintf("(%s)", ifs.Condition))
	parsedExpression, _, errs := MathExpressionTokensToEnd(lx.NextToken(), lx)
	if errs != nil {
		return false, errs
	}
	if len(parsedExpression) == 2 {
		return false, nil
	}
	val, e := ParseArithmeticExpressions(parsedExpression[1:len(parsedExpression)-1], lx)
	if e != nil {
		return false, e
	}
	if v, ok := val.(bool); ok {
		return v, nil
	}
	return false, err.NewInvalidIfExpressionError(lx.Scanner.Pos())
}
