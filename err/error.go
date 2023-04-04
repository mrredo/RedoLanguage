package err

import (
	"fmt"
	"text/scanner"
)

func NewError(msg string, pos scanner.Position) error {
	return fmt.Errorf("%s:%d:%d: %s", pos.Filename, pos.Line, pos.Column, msg)
}
