package ast

import "RedoLanguagev2/types"

func ParseTokens(tokens []types.Token) (types.AST, []error) {
	errors := []error{}
	var ast = types.AST{}
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		switch v.Type {
		case types.Var:
			varT := types.Node{Type: types.IdentifierDeclaration, Body: []types.Node{}}
			i++
			key := tokens[i]

			varT.Body = append(varT.Body, types.Node{
				Type:  types.Identifier,
				Value: key.Value,
			})
			i++
			op := tokens[i]
			varT.Body = append(varT.Body, types.Node{
				Type:  op.Type,
				Value: key.Value,
			})
			_ = op
			i++
			expression := types.Node{Type: types.Expression}
			_ = expression
			//expression
			for i < len(tokens) {
				tok := tokens[i]

				if tok.Type == types.NewLine || tok.Type == types.SemiColon {
					break
				}
				expression.Value += tok.Value
				i++
			}
			ast.Body = append(ast.Body, varT)
		}
	}

	return ast
}
