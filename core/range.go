package main

import "fmt"

/*
** Range iterates over elements in a variety of data structures
 */

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// When using range on arrays and slices it return both the key (index) and value
	// Here since with did not need the index we ignore it using _ identifier
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, v := range nums {
		fmt.Println(i, "->", v)
	}

	// Print the index which hold the value 3
	for i, num := range nums {
		if num == 3 {
			fmt.Println("3 is stored at index:", i)
		}
	}

	// Range over maps iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	// Iterate over map keys only
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range on strings iterates over Unicode code points (characters). The
	// first value (i) is the index and the second (c) is the value.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
