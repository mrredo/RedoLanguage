package types
type AST struct {
	Blocks []TokenAST
}
type TokenAST interface {
	IsToken()
}
type GeneralToken struct {
	
}
func (t *GeneralToken) IsToken() {}
type IdentifierFunction struct {
	Key string
}

func (t *IdentifierToken) IsToken() {}
type IdentifierToken struct {
	Type TokenType
	Key string
}
func (t *IdentifierFunction) IsToken() {}
type ExpressionAST struct {
	Expression string
}
func (t *ExpressionAST) IsToken() {}