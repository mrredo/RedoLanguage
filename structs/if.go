package structs


type If struct {
	Condition
	Position int

	Output bool
}
type ElseIf struct {
	Position int
	Condition string
	Output bool
}
type Else struct {
	Position int
	ConditionString string // has condition since it 
	Output bool
}
type ConditionEvaluator interface {
	EvaluateCondition() bool
}

type Condition struct {
	Condition string
	ConditionEvaluator // embed the interface in a new struct
}
func (c Condition) EvaluateCondition() bool {
	// Evaluate the condition associated with the statement
	// and return its boolean value
	return evaluateCondition(c.Condition)
}
func testst() {
	ifs := If{
		Condition: Condition{
			Condition: "true",
		},
		Position: 1,
		Output: true,
	}
	ifs.Condition.EvaluateCondition()
}
func evaluateCondition(condition string) bool {
	// Evaluate the condition and return its boolean value
	// You can implement this function based on the syntax and semantics of your programming language
	return true
}