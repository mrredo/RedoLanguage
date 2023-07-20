package types

import gtoken "go/token"

type TokenType int

const (
	Number TokenType = iota
	Plus
	PlusAssign
	Minus
	MinusAssign
	Multiply
	MultAssign
	Divide
	DivAssign
	LeftParenthesis
	RightParenthesis
	NewLine
	SemiColon
	String
	Identifier
	Var
	If
	LBrack
	RBrack
	Unknown

	//AST
	IdentifierDeclaration
	Expression
	Assign
	AssignExpression
	IdentifierExpression
)

type Token struct {
	Type  TokenType
	Value string
	Pos   gtoken.Position
}
