package main

import "fmt"

/*
** Go supports methods defined on structs
 */

type rect struct {
	width, height int
}

// The area method has a receiver type of *react
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Call the struct methods
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// Go automatically handles conversion between values and pointers for
	// method calls.
	// Using pointer receiver types is useful to avoid copying on method calls
	// or to allow the method to mutate the receiving struct
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
