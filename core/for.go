package main

import "fmt"

func main() {
	// Basic 'for' usage
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// The classic 'for' with intial, condition and increment
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Indefinite 'for'. Only way to get out is to 'break' out of it or return
	// from the enclosing function
	for {
		fmt.Println("single loop")
		break
	}

	// 'continue' with the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
