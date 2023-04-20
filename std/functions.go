package std

import (
	"fmt"
)

var Functions = map[string]func(args ...interface{}) interface{}{
	"print": func(args ...interface{}) interface{} {
		fmt.Print(args...)
		return nil
	},

	"printf": func(args ...interface{}) interface{} {
		if len(args) == 0 {
			return ""
		}
		arg1, ok1 := args[0].(string)
		if !ok1 {
			fmt.Println("printf function requires first argument to be string")
			return nil
		}
		if len(args) == 1 {
			return arg1
		}
		fmt.Printf(arg1, args[1:]...)
		return nil
	},
	"println": func(args ...interface{}) interface{} {
		fmt.Println(args...)
		return nil
	},

	"free": func(args ...interface{}) interface{} {
		if len(args) > 1 {
			fmt.Println("free function takes 1 argument")
			return nil
		}
		str, ok := args[0].(string)
		if !ok {
			fmt.Println("free function first argument must be a string")
			return nil
		}
		delete(Variables, str)
		return nil
	},
}
