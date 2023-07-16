package ast

import "RedoLanguagev2/types"


func AssignExpression(ident types.Token, op types.Token, expression []types.Token) (types.Node) {
	ast := types.Node{Type: types.AssignExpression, Name: ident.Value, Value: op.Value}
	ast.Expression = &types.Node{
		Type: types.Expression,
		Value: "",
	}
	for _, v := range expression {
		ast.Expression.Value += v.Value
	}
	return ast
}