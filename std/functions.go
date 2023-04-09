package std

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

var Functions = map[string]func(args ...interface{}) interface{}{
	"print": func(args ...interface{}) interface{} {
		fmt.Print(args...)
		return nil
	},
	"math": func(args ...interface{}) interface{} {
		if len(args) > 1 {
			fmt.Println("math function takes 1 argument")
			return nil
		}
		str, ok := args[0].(string)
		if !ok {
			fmt.Println("math function first argument must be a string")
			return nil
		}
		output, err := ParseArithmeticExpressions(str)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return output
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

func ParseArithmeticExpressions(expression string) (any, error) {
	// Create new expression with default token factory
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, fmt.Errorf("Error parsing expression: %v", err)
	}

	// Evaluate expression with empty parameter map
	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, fmt.Errorf("Error evaluating expression: %v", err)
	}

	// Convert result to int and return

	if result == true || result == false {
		return result == "true", nil
	}
	if val, ok := result.(float64); ok {
		return int(val), nil
	}
	return 0, fmt.Errorf("error converting result to int")
}
