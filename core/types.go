package main

import (
	"fmt"
	"reflect"
)

var x interface{} = []int{1, 2, 3}
var y interface{} = 2.3

func main() {
	idByPrintSpecifiers()
	idBySwitch()
	idByReflection()
}

func idByPrintSpecifiers() {
	t := fmt.Sprintf("%T", x)
	fmt.Println("Type of x: " + t)
}

func idBySwitch() {
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

func idByReflection() {
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType, xValue)
}
