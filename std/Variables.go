package std

var Variables = map[string]Variable{
	// "e": {
	// 	Key: "e",
	// 	Value: 10,
	// },
}

type Variable struct {
	Value        any
	Key          string
	NestingLevel int
	ID           int
}

func (this *Variable) SetValue(val any) {
	this.Value = val
}
