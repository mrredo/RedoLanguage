package std

import "fmt"

var Functions = map[string]func(args ...interface{}) interface{}{
	"print": func(args ...interface{}) interface{} {
		fmt.Print(args...)
		return nil
	},
	"println": func(args ...interface{}) interface{} {
		fmt.Println(args...)
		return nil
	},
	"modulo": func(args ...interface{}) interface{} {
		if len(args) != 2 {
			fmt.Println("multiply function takes exactly two arguments")
			return nil
		}
		arg1, ok1 := args[0].(int)
		arg2, ok2 := args[1].(int)
		if !ok1 || !ok2 {
			fmt.Println("multiply function arguments must be numbers")
			return nil
		}
		return arg1 % arg2
	},
	"subtract": func(args ...interface{}) interface{} {
		if len(args) != 2 {
			fmt.Println("multiply function takes exactly two arguments")
			return nil
		}
		arg1, ok1 := args[0].(int)
		arg2, ok2 := args[1].(int)
		if !ok1 || !ok2 {
			fmt.Println("multiply function arguments must be numbers")
			return nil
		}
		return arg1 - arg2
	},
	"multiply": func(args ...interface{}) interface{} {
		if len(args) != 2 {
			fmt.Println("multiply function takes exactly two arguments")
			return nil
		}
		arg1, ok1 := args[0].(int)
		arg2, ok2 := args[1].(int)
		if !ok1 || !ok2 {
			fmt.Println("multiply function arguments must be numbers")
			return nil
		}
		return arg1 * arg2
	},
	"divide": func(args ...interface{}) interface{} {
		if len(args) != 2 {
			fmt.Println("divide function takes exactly two arguments")
			return nil
		}
		arg1, ok1 := args[0].(int)
		arg2, ok2 := args[1].(int)
		if !ok1 || !ok2 {
			fmt.Println("divide function arguments must be numbers")
			return nil
		}
		return arg1 / arg2
	},
	"add": func(args ...interface{}) interface{} {
		if len(args) != 2 {
			fmt.Println("add function takes exactly two arguments")
			return nil
		}
		arg1, ok1 := args[0].(int)
		arg2, ok2 := args[1].(int)
		if !ok1 || !ok2 {
			fmt.Println("add function arguments must be numbers")
			return nil
		}

		return arg1 + arg2
	},
}
