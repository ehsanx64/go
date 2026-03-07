package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	// Custom functions
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...any) (any, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
	}

	// Expression
	exp := "(number_1 > 0 && number_2 > 1) || (strlen(text_3) > 15)"

	// Parameters
	parameters := make(map[string]any, 8)
	parameters["number_1"] = -100
	parameters["number_2"] = 100
	parameters["text_3"] = "Adam & Eve"

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(exp, functions)
	if err != nil {
		fmt.Println("Error prepating the expression for evaluation")
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println("Error evaluating the expression")
	}

	fmt.Printf("%s\n", exp)
	fmt.Printf("%s: %d\n", "number_1", parameters["number_1"])
	fmt.Printf("%s: %d\n", "number_2", parameters["number_2"])
	fmt.Printf("%s: %s\n", "text_3", parameters["text_3"])
	fmt.Printf("Result: %+T\t %+v\n", result, result)
}
