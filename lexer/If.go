package lexer

import "RedoLanguage/err"

type If struct {
	Condition    string
	Position     int
	NestingLevel int
	//ElseIfs []elseIf
}

//	type elseIf struct {
//		Condition string
//
// }
//
//	func (ifs *elseIf) Output(lx *Lexer) (bool, error) {
//		val, e := ParseArithmeticExpressions(ifs.Condition, lx)
//		if e != nil {
//			return false, e
//		}
//		if v, ok := val.(bool); ok {
//			return v, nil
//		}
//		return false, err.NewInvalidIfExpressionError(lx.Scanner.Pos())
//	}
func (ifs *If) Output(lx *Lexer) (bool, error) {
	val, e := ParseArithmeticExpressions(ifs.Condition, lx)
	if e != nil {
		return false, e
	}
	if v, ok := val.(bool); ok {
		return v, nil
	}
	return false, err.NewInvalidIfExpressionError(lx.Scanner.Pos())
}
