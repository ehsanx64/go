package main

import "fmt"

var a = 6

func main() {
	// Since we are using a as a local variable, mutating it will not affect the
	// a variable in global scope
	fmt.Println("Using local variable...")
	print()
	demoLocal()
	print()

	// Using the a variable in global scope context mutates it
	fmt.Println("Using global variable...")
	print()
	demoGlobal()
	print()
}

// Just print the a variable
func print() {
	fmt.Println(a)
}

/*
** Following two functions are almost the same, but introduction of := instead
** of = in demoLocal function identifies the a as a local variable within the
** function so the assignment won't affect the a variable in global scope.
**
** It can be a source of errors, so watch out!!!
 */

func demoLocal() {
	a := 5
	fmt.Println(a)
}

func demoGlobal() {
	a = 5
	fmt.Println(a)
}
