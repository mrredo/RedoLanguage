package types

type AST struct {
	Body []Node
}

type Node struct {
	Type  TokenType
	Value string
	Name  string
	//callee     *Node
	Expression *Node
	Body       []Node
	//params     []Node
	Arguments *[]Node
	//context    *[]Node
}
