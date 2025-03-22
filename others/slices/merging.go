package main

import "fmt"

/*
 * Merge two slices with duplicates removed
 */
func mergeSlicesWithoutDuplicates(slice1, slice2 []int) []int {
	// Create a map to track unique elements
	unique := make(map[int]bool)
	result := []int{}

	// Add elements from the first slice
	for _, v := range slice1 {
		if !unique[v] {
			unique[v] = true
			result = append(result, v)
		}
	}

	// Add elements from the second slice
	for _, v := range slice2 {
		if !unique[v] {
			unique[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}

	merged := mergeSlicesWithoutDuplicates(slice1, slice2)
	fmt.Println(merged) // Output: [1 2 3 4 5 6]
}

