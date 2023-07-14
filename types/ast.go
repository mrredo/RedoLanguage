package types

type AST struct {
	Blocks []TokenAST
}
type TokenAST interface {
	IsToken()
}
type IdentifierFunction struct {
	Key string
}

func (t *IdentifierToken) IsToken() {}
type IdentifierToken struct {
	Type TokenType
	Key string
}
func (t *IdentifierFunction) IsToken() {}