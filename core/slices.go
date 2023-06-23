package main

import "fmt"

/*
** Slices can be viewed as an imporved version of arrays. Unlike arrays, slices
** are only typed by the elements they contain (not the number of elements). An
** uninitialized slice equals to nil and has 0 elements.
 */

func main() {
	// Make a slice of string with length of 3
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Put some values into the slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// Append a new element
	s = append(s, "d")
	// Append multiple elements
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Copy a slice to another (newly declared) slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// ':' is the slice operator and can be used to slice a slice.
	// 'l' would be third-element to last-element of the 's' (including 2 and
	// excluding 5)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Slice frow first to end of slice (0 to 5 - excluding 5)
	l = s[:5]
	fmt.Println("sl2:", l)

	// Slice from third element to end of slice (including 2)
	l = s[2:]
	fmt.Println("sl3:", l)

	// Define and initialize a slice in the same line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional structures like arrays.
	// Unlike arrays the inner slices length can vary.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
