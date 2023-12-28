package main

import (
	"encoding/json"
	"fmt"
	"math/cmplx"
	"reflect"
	"strings"
)

var x interface{} = []int{1, 2, 3}
var y interface{} = 2.3

const tableWidth = 17

func main() {
	// Identify types using printf specifiers
	idByPrintSpecifiers()

	// Identify types using switch
	idBySwitch()

	// Identify types by reflection
	idByReflection()

	// Simple types demo
	demoTypes()

	// Type inference demo
	demoTypeInference()

	// A simple showcase of the any type
	demoAnyType()

	// Using the any type for unmarshalling dynamic (json) data
	demoAnyAndUnmarshall()
}

func demoAnyType() {
	fmt.Println()
	fmt.Println("* Any type demo")

	// Following line is same as var anyOne interface{}
	var anyOne any

	fmt.Println("** Default value")
	fmt.Printf("Type:  %+T\n", anyOne)
	fmt.Printf("Value: %+v\n\n", anyOne)

	fmt.Println("** After int assignment")
	anyOne = 10
	fmt.Printf("Type:  %+T\n", anyOne)
	fmt.Printf("Value: %+v\n\n", anyOne)

	fmt.Println("** After string assignment")
	anyOne = "text"
	fmt.Printf("Type:  %+T\n", anyOne)
	fmt.Printf("Value: %+v\n\n", anyOne)

	fmt.Println("** After map assignment")
	anyOne = make(map[string]int)
	fmt.Printf("Type:  %+T\n", anyOne)
	fmt.Printf("Value: %+v\n\n", anyOne)

	fmt.Println("** After slice assignment")
	anyOne = make([]int, 5)
	fmt.Printf("Type:  %+T\n", anyOne)
	fmt.Printf("Value: %+v\n\n", anyOne)

	fmt.Println("** After array assignment")
	anyOne = [5]int{0, 1, 2, 3, 4}
	fmt.Printf("Type:  %+T\n", anyOne)
	fmt.Printf("Value: %+v\n\n", anyOne)
}

func dumpValue(title string, arg any) {
	var padding strings.Builder

	for i := 0; i < tableWidth-(len(title)+1); i++ {
		padding.WriteString(" ")
	}

	fmt.Printf("%s:%s%+v\n", title, padding.String(), arg)
}

func dumpType(title string, arg any) {
	var padding strings.Builder

	for i := 0; i < tableWidth-(len(title)+1); i++ {
		padding.WriteString(" ")
	}

	fmt.Printf("%s:%s%+T\n", title, padding.String(), arg)
}

/*
** A basic use case for the any type is when unmarshalling json data which are
** dynamic and cannot be predicted.
 */
func demoAnyAndUnmarshall() {
	fmt.Println()
	fmt.Println("* Any and unmarshall demo")

	// Following line is same as var anyOne interface{}
	data := `{"name": "Adam", "age": 33, "skills": ["dreaming", "hunting"], "weapons": ["bag", "pen"]}`

	fmt.Println("** Default json data")
	dumpValue("Value", data)
	dumpType("Type", data)
	fmt.Println()

	var obj any
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("** After unmarshalled data (object)")
	dumpValue("Value", obj)
	dumpType("Type", obj)
	fmt.Println()

	fmt.Println("** Get values as object properties")
	dumpValue("Name", obj.(map[string]any)["name"])
	// or a shorter and better way
	u := obj.(map[string]any)
	dumpValue("Age", u["age"])
	dumpValue("Weapons", u["weapons"])

	skills := u["skills"].([]any)
	dumpValue("Skills", skills)
	dumpValue("First skill", skills[0])
	dumpValue("Second skill", skills[1])
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
	fmt.Println()
	fmt.Println("* idByReflection()")
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType, xValue)
	fmt.Println("Kind:", xValue.Kind())
}
