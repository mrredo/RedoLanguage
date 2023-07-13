package types

type AST struct {
	//imports []ImportToken
	Blocks []ASTtoken
}
var As = AST{
	Blocks: []ASTtoken{
		VariableDeclaration{
			
		},
	},
}
type ASTtoken interface {
	isToken()
}
type Expression struct {

}
type VariableDeclaration struct {
	// Variable declaration fields
	Type TokenType
	Key string
	
	Value  Expression
}

func (vd VariableDeclaration) isToken() {}