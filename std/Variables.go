package std

var Variables = map[string]Variable{}

type Variable struct {
	Value        any
	Key          string
	NestingLevel int
	ID           int
}

func (this *Variable) SetValue(val any) {
	this.Value = val
}
