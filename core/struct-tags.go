package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Animal struct {
	Name string `required max_len:"2"`
	Age  int    `Age of the animal`
}

type T struct {
	F1 string `json:"field1"`
	F2 string `json:"field2, omitempty"`
	F3 string `json:"field3, omitempty"`
	F4 string `json:"-"`
}

func simpleDemo() {
	a := Animal{"Cat", 10}
	t := reflect.TypeOf(a)
	field, _ := t.FieldByName("Age")

	fmt.Printf("Type = %T, Value = %v\n", a, a)
	fmt.Printf("Type = %T, Value = %v\n", t, t)
	fmt.Printf("Type = %T, Value = %v\n", field.Tag, field.Tag)
}

func jsonDemo() {
	t := T{
		F1: "v1",
		F2: "",
		F3: "v3",
		F4: "v4",
	}

	s, _ := json.Marshal(t)
	fmt.Printf("JSON: %s\n", s)
}

func main() {
	simpleDemo()
	jsonDemo()
}
