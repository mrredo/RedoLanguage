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
	
	Unknown
)

type Token struct {
	Type  TokenType
	Value string
}