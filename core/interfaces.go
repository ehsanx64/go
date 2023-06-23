package main

import (
	"fmt"
	"math"
)

/*
** Interfaces are named collections of method signatures.
 */

// Define a basic interface
type geometry interface {
	area() float64
	perim() float64
}

// Define two different sturcts (rect, circle)
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Define all the methods defined in the interface, that our structs are going
// to use
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, it will be used to lookup for its methods
// The measure function can operate on any struct that implements the geometry
// interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Use same function for different structs
	measure(r)
	measure(c)
}
