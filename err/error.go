package err

import (
	"fmt"
	"text/scanner"
)

func Error(errFormat string, errType string, pos scanner.Position) error {
	return fmt.Errorf("./%s:%d:%d: %s: %s", pos.Filename, pos.Line, pos.Column, errType, errFormat)
}
func NewUnusedError(variableName string, pos scanner.Position) error {
	return Error(fmt.Sprintf(VariableNotUsed, variableName), RefrenceError, pos)
}
func NewUndefinedError(variableName string, pos scanner.Position) error {
	return Error(fmt.Sprintf(VariableNotDefined, variableName), RefrenceError, pos)
}
func NewTypeError(pos scanner.Position) error {
	return Error(MismatchedTypesInExpression, TypeError, pos)
}
func NewExpressionError(format string, msg string, pos scanner.Position) error {
	return Error(fmt.Sprintf(format, msg), ExpressionError, pos)
}
func NewFunctionVoidError(variableName string, pos scanner.Position) error {
	return Error(fmt.Sprintf(FunctionReturnsVoid, variableName), InvalidFunctionCall, pos)
}
func NewSyntaxError(format string, pos scanner.Position, args ...any) error {
	return Error(fmt.Sprintf(format, args), SyntaxError, pos)
}
func NewSemiColonError(pos scanner.Position) error {
	return Error(MissingSemicolon, SyntaxError, pos)
}
func NewIllegalTokenError(tokenVal string, pos scanner.Position) error {
	return Error(fmt.Sprintf(IllegalToken, tokenVal), SyntaxError, pos)
}
