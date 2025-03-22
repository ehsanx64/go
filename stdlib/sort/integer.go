package main

import (
	"fmt"
	"sort"
)

func main() {
	uints := []uint{5, 2, 6, 3, 1, 4}

	ints := make([]int, 0, len(uints))
	for i := range uints {
		ints = append(ints, int(uints[i]))
	}

	fmt.Println("Initial:")
	fmt.Println(uints)

	sort.Ints(ints)
	fmt.Println("Sorted:")
	fmt.Println(ints)
}
