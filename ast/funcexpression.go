package ast

import "RedoLanguagev2/types"

func IdentExpression(name types.Token, args []types.Token) types.Node {
	fnExp := types.Node{
		Type: types.IdentifierExpression,
		Name: name.Value,
	}
}
