package lexer

func ConvertToTokenType(typeS string) TokenType {
	switch typeS {
	case "string":
		return STRING
	case "int", "int64", "int32", "int16", "int8":
		return NUMBER
	case "bool", "boolean":
		return BOOL

	default:
		return IDENTIFIER
	}
}
