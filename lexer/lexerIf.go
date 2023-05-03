package lexer

import (
	"errors"
)

func ExecuteIf(c, s Token, lx *Lexer) error {
		curNes := lx.CurrentNestingLevel
		tok := s
		i := If{Position: lx.PositionMap[lx.CurrentNestingLevel], NestingLevel: lx.CurrentNestingLevel}
		
		//i := If{Position: lx.CurrentPosition, NestingLevel: lx.CurrentNestingLevel}
		forif:
		for {
			switch tok.Type {
			case EOF:
				return nil
			case LBRACE:
				
				lx.CurrentPosition++
				lx.CurrentNestingLevel++
				
				break forif
			}
		
			i.Condition += tok.Value
			tok = lx.NextToken()
		}
		if curNes == lx.CurrentNestingLevel {
			return errors.New("missing start of statement")
		}
		curNes = lx.CurrentNestingLevel
		lx.IfPositions[lx.CurrentNestingLevel] = i
		lx.PositionMap[lx.CurrentNestingLevel]++
		if v, err := i.Output(); err != nil {
			return err
		} else {
			if !v {
			forfalse:
				for {
				
					switch tok.Type {
					case EOF:
						break forfalse
					case RBRACE:
					
						lx.CurrentNestingLevel--
						if lx.CurrentNestingLevel == curNes {
							break forfalse
						}
					
					case LBRACE:
						lx.CurrentNestingLevel++
					}
				
					tok = lx.NextToken()
				
				}
			
			}
		}
		return nil
}
func IsIfStatement(c Token) bool {
	return c.Type == IF || c.Type == ELSE || c.Type == ELSE_IF
}
