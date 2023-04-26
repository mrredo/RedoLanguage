package interpreter

import "testing"

func TestInterpret(t *testing.T) {
	type args struct {
		input    string
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test interpreter 1",
			args: args{
				input: `
				
				if true {
					println(true)
				}
				`,
				fileName: "test1.rd",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interpret(tt.args.input, tt.args.fileName)
		})
	}
}
