package main

import "fmt"

func demoBasic() {
	// Define an array of ints with 5 elements. As default value for int's is 0,
	// this array is an array of 5 zero's. Type of elements and length of an
	// array are part of its type.
	var a [5]int

	// Println will dump entire array
	fmt.Println("a:", a)

	// Array are zero-indexed, so following code puts 100 in the fifth (last)
	// position
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin len function returns the length of an array
	fmt.Println("len:", len(a))

	// Using braces to declare and initialize the array (in one go)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
}

func demo2dArray() {
	// Define a two-dimensional array.
	// Array types are one-dimensional but programmers can compose types to
	// build multi-dimensional data structures
	var twoD [2][3]int

	// Using two nested loops set values for twoD array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	// Arrays are wrapped in [ and ] when printed
	fmt.Println("2d:", twoD)
}

func demoDotsOperator() {
	type Item struct {
		name  string
		price int
	}

	// Using the '...' operator to set the length of array automatically
	var items = [...]Item{
		{name: "Laptop", price: 1000},
		{name: "Tablet", price: 700},
		{name: "Cellphone", price: 500},
	}

	// Dump the array
	fmt.Println(items)
}

func main() {
	demoBasic()
	demo2dArray()
	demoDotsOperator()
}
