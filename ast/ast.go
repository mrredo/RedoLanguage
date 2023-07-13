package ast

type AST struct {
	//imports []ImportToken
	Blocks []ASTtoken
}
type ASTtoken interface {
	isToken()
}
type VariableDeclaration struct {
	// Variable declaration fields
	Type 
}

func (vd VariableDeclaration) isToken() {}