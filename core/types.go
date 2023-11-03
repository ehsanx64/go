package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
)

var x interface{} = []int{1, 2, 3}
var y interface{} = 2.3

func main() {
	idByPrintSpecifiers()
	idBySwitch()
	idByReflection()
	demoTypes()
	demoTypeInference()
}

func demoTypeInference() {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i

	fmt.Println()
	fmt.Println("* Type inference demo")
	fmt.Printf("i type %T, value %v\n", i, i)
	fmt.Printf("j type %T, value %v\n", j, j)
	fmt.Printf("k type %T, value %v\n", k, k)
}

func demoTypes() {
	var (
		ToBe   bool       = false
		anInt  int        = 39
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Println()
	fmt.Println("* Types demo")
	fmt.Printf("Type: %T\t\tValue: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T\t\tValue: %v\n", anInt, anInt)
	fmt.Printf("Type: %T\t\tValue: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T\tValue: %v\n", z, z)
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
