package tokens

type TokenType int

const (
	EOF    TokenType = iota // end of file
	LPAREN                  // (
	RPAREN                  // )
	LBRACE                  // {
	RBRACE                  // }

	COMMA                 // ,
	IDENTIFIER            // variable identifier
	NUMBER                // numeric literal
	STRING                // string literal
	EQUAL                 // ==
	ASSIGN                // =
	VAR                   // var keyword
	BOOL                  // boolean literal
	MODULO                // %
	DIVIDE                // /
	PLUS                  // +
	MULTIPLY              // *
	SUBTRACT              // -
	PLUS_ASSIGN           // +=
	SUBTRACT_ASSIGN       // -=
	MULTIPLY_ASSIGN       // *=
	DIVIDE_ASSIGN         // /=
	MODULO_ASSIGN         // %=
	LEFT_SHIFT_ASSIGN     // <<=
	RIGHT_SHIFT_ASSIGN    // >>=
	BITWISE_AND_ASSIGN    // &=
	BITWISE_XOR_ASSIGN    // ^=
	PLUS_PLUS             // ++
	SUBTRACT_SUBTRACT     // --
	BITWISE_XOR           // ^
	LEFT_SHIFT            // <<
	RIGHT_SHIFT           // >>
	BITWISE_AND           // &
	EQUAL_TO              // ==
	NOT_EQUAL_TO          // !=
	LESS_THAN             // <
	LESS_THAN_OR_EQUAL    // <=
	GREATER_THAN          // >
	GREATER_THAN_TO_EQUAL // >=
	AND                   // &&
	SEMICOLON             // ;
	NEW_LINE              // \n
	BITWISE_OR            // |
	OR                    // ||
	ILLEGAL

	IF      // if
	ELSE    // else
	ELSE_IF // else if

)
var keywords = map[string]TokenType{
	"if": IF,
	"else if": ELSE_IF,
	"else": ELSE,
}