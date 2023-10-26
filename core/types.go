package main

import "fmt"

func main() {
	var x interface{} = []int{1, 2, 3}
	var y interface{} = 2.3

	t := fmt.Sprintf("%T", x)
	fmt.Println("Type of x: " + t)

	fmt.Print("Type of y: ")

	switch y.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
