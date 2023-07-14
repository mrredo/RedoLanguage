package types

type TokenType int

const (
	Number TokenType = iota
	Plus
	Minus
	Multiply
	Divide
	LeftParenthesis
	RightParenthesis
	NewLine
	String
	Identifier
	Var
	Unknown
)

type Token struct {
	Type  TokenType
	Value string
}
