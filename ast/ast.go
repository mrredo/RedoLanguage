package ast

import "RedoLanguagev2/types"

func ParseTokens(tokens []types.Token) types.AST {
	var ast = types.AST{}
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		switch v.Type {
		case types.Var:
			i++
			key := tokens[i]
			_ = key
			i++
			op := tokens[i]
			_ = op
			i++
			expression := types.ExpressionAST{}
			 _ = expression
			//expression
			for i < len(tokens) {
				tok := tokens[i]
				
				if tok.Type == types.NewLine || tok.Type == types.SemiColon {
					break
				}
				expression.Expression += tok.Value
				i++
			}
		}
	}
	switch {

	}
	return ast
}
