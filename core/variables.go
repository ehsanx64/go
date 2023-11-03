package main

import "fmt"

// Demonstrate type based (automatic) default value assignment
func demoDefaultValues() {
	// Define some variables with no initial values
	var (
		boolVal bool
		intVal  int
	)

	fmt.Println("boolVal:", boolVal)
	fmt.Println("intVal:", intVal)

	var javascript, python, java bool
	var i, j int
	var k1 rune
	var k2 rune = 'b'
	var s1, s2, s3 string = "", "*", "..."

	fmt.Println("javascript, python, java =>", javascript, python, java)
	fmt.Println("i, j =>", i, j)
	fmt.Println("k1, k2 =>", k1, k2)
	fmt.Println("s1=\"\", s1=\"*\", s1=\"...\" =>", s1, s2, s3)
}

func main() {
	// Define and initialize a variable so the 'var' can infer its type
	var a = "initial"
	fmt.Println(a)

	// := operator defines a variable and infer its type by its initial value
	f := "apple"
	fmt.Println(f)

	// Declare multiple variables at once and initialize them (both are ints)
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Using type inference to declare and initialize a boolean value
	var d = true
	fmt.Println(d)

	// Declare variable 'e', but since it is not initialized, it would have the
	// default initialization value which for int type is 0
	var e int
	fmt.Println("e =>", e)

	demoDefaultValues()
}
