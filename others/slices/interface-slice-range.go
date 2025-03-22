package main

import "fmt"

func main() {
	target := []interface{}{
		"Adam",
		"Eve",
	}

	fmt.Println("Target:", target)
	fmt.Printf("%T: %v\n", target, target)

	// Range over it
	for i, v := range target {
		fmt.Printf("%02d: %s\n", i, v)
	}
}
