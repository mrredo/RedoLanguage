package types

type AST struct {
	//imports []ImportToken
	Blocks []ASTtoken
}
type ASTtoken interface {
	isToken()
}
type Expression struct {

}
type VariableDeclaration struct {
	// Variable declaration fields
	Type TokenType
	key string
	value  Expression
}

func (vd VariableDeclaration) isToken() {}