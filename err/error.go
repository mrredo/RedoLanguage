package err

func Error() {

}

//	func Error(errFormat string, errType string, pos scanner.Position) error {
//		return fmt.Errorf("./%s:%d:%d: %s: %s", pos.Filename, pos.Line, pos.Column, errType, errFormat)
//	}
//
//	func NewUnusedError(variableName string, pos scanner.Position) error {
//		return Error(fmt.Sprintf(VariableNotUsed, variableName), RefrenceError, pos)
//	}
//
//	func NewUndefinedError(variableName string, pos scanner.Position) error {
//		return Error(fmt.Sprintf(VariableNotDefined, variableName), RefrenceError, pos)
//	}
//
//	func NewTypeError(pos scanner.Position) error {
//		return Error(MismatchedTypesInExpression, TypeError, pos)
//	}
//
//	func NewExpressionError(format string, msg string, pos scanner.Position) error {
//		return Error(fmt.Sprintf(format, msg), ExpressionError, pos)
//	}
//
//	func NewFunctionVoidError(variableName string, pos scanner.Position) error {
//		return Error(fmt.Sprintf(FunctionReturnsVoid, variableName), InvalidFunctionCall, pos)
//	}
//
//	func NewSyntaxError(format string, pos scanner.Position, args ...any) error {
//		return Error(fmt.Sprintf(format, args), SyntaxError, pos)
//	}
//
//	func NewSemiColonError(pos scanner.Position) error {
//		return Error(MissingSemicolon, SyntaxError, pos)
//	}
//
//	func NewInvalidIfExpressionError(pos scanner.Position) error {
//		return Error(IfExpressionResultNotBool, ExpressionError, pos)
//	}
//
//	func NewIllegalTokenError(tokenVal string, pos scanner.Position) error {
//		return Error(fmt.Sprintf(IllegalToken, tokenVal), SyntaxError, pos)
//	}
const (
	VariableNotDefined          = "'%s' is not defined"
	VariableNotUsed             = "'%s' is not used"
	MismatchedTypesInExpression = "mismatched types in expression"
	FunctionReturnsVoid         = "'%s' method returns void"
	VariableAlreadyDeclared     = "'%s' is already declared"
	//math
	ErrorParsingExpression     = "failed parsing math expression - %s"
	ErrorEvaluatingExpression  = "failed evaluating math expression - %s"
	ErrorConvertingResultToInt = ""
	UnbalancedParentheses      = "unbalanced parentheses"
	//
	MissingSemicolon = "missing semicolon"
	IllegalToken     = "'%s' is illegal token"

	IfExpressionResultNotBool = "invalid expression in if statements expression must return a boolean"
)
