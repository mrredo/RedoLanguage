package err

import (
	"fmt"
	"text/scanner"
)

func NewError(msg string, pos scanner.Position) error {
	return fmt.Errorf("%s:%d:%d: %s", pos.Filename, pos.Line, pos.Column, msg)
}
func Error(errFormat string, errType string, pos scanner.Position) error {
	return fmt.Errorf("./%s:%d:%d: %s: %s", pos.Filename, pos.Line, pos.Column, errType, errFormat)
}
func NewUndefinedError(variableName string, pos scanner.Position) error {
	return Error(fmt.Sprintf(VariableNotDefined, variableName), RefrenceError, pos)
}
func NewTypeError(pos scanner.Position) error {
	return Error(MismatchedTypesInExpression, TypeError, pos)
}
