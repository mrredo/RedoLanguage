package err

const (
	VariableNotDefined          = "'%s' is not defined"
	VariableNotUsed             = "'%s' is not used"
	MismatchedTypesInExpression = "mismatched types in expression"
	FunctionReturnsVoid         = "'%s' method returns void"
	VariableAlreadyDeclared     = "'%s' is already declared"
	//math
	ErrorParsingExpression     = "failed parsing math expression - %s"
	ErrorEvaluatingExpression  = "failed evaluating math expression - %s"
	ErrorConvertingResultToInt = ""
	UnbalancedParentheses      = "unbalanced parentheses"
	//
	MissingSemicolon = "missing semicolon"
)
