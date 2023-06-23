package main

import "fmt"

// Factorial function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("7!:", fact(7))

	// Closures can be recursive, but this requires the closure to be declared
	// with a typed var explicitly before it is defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println("fib(7):", fib(7))
}
