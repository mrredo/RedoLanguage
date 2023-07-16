package ast

import (
	"RedoLanguagev2/types"
	"errors"
)

func ParseTokens(tokens []types.Token) (types.AST, []error) {
	errorl := []error{}
	var ast = types.AST{}
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		switch v.Type {
		case types.Identifier:
			i++
			sec := tokens[i]
			switch sec.Type {
			case types.DivAssign, types.PlusAssign, types.MinusAssign, types.MultAssign:
				expression := []types.Token{}
				i++
				for i < len(tokens) {
					tok := tokens[i]
	
					if tok.Type == types.NewLine || tok.Type == types.SemiColon {
						break
					}
					expression = append(expression, tok)
					i++
				}
				astAss := AssignExpression(v, sec, expression)
				ast.Body = append(ast.Body, astAss)
			} 
		case types.Var:
			varT := types.Node{Type: types.IdentifierDeclaration, Body: []types.Node{}}
			i++
			key := tokens[i]
			if key.Type != types.Identifier {

				errorl = append(errorl, errors.New("invalid type for key"))
			}
			i++
			op := tokens[i]

			if op.Type != types.Assign {

				errorl = append(errorl, errors.New("invalid operator for identifier declaration"))
			}
			i++

			expression := &types.Node{Type: types.Expression}
			//expression
			for i < len(tokens) {
				tok := tokens[i]

				if tok.Type == types.NewLine || tok.Type == types.SemiColon {
					break
				}
				expression.Value += tok.Value
				i++
			}
			varT.Name = key.Value
			varT.Expression = expression
			ast.Body = append(ast.Body, varT)
		}
	}

	return ast, errorl
}
