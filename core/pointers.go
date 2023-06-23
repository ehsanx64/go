package main

import "fmt"

/*
** Pointers are used to pass references to value and strucutres
 */

// This function recieves the parameter as a value, so it gets a copy of it, and
// operations on the variable (ival) has no effect outside of this function
func zeroval(ival int) {
	ival = 0
}

// This function recieves the parameter as a reference (pointer) to an int.
func zeroptr(iptr *int) {
	// Dereference the pointer and assign a value to it, which ultimately changes
	// the parameter value
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// Calling zeroval has no effect on i (passed by value)
	zeroval(i)
	fmt.Println("zeroval:", i)

	// Calling zeroptr with effect (passed by reference)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// & operator returns address of a variable
	fmt.Println("pointer", &i)

}
