package main

import "fmt"

func main() {
	fmt.Println("* Basic 'for' usage")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()

	fmt.Println("* The classic 'for' with intial, condition and increment")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println()

	// Only way to get out is to 'break' out of it or return from the enclosing
	// function
	fmt.Println("* Indefinite 'for'")
	for {
		fmt.Println("single loop")
		break
	}
	fmt.Println()

	fmt.Println("* 'continue' with the next iteration")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
