package main

import "fmt"

func main() {
	// Simple if, without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if-else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if-elseif-else
	// The num variable is defined and initalized immediately before the if's
	// first condition check
	// This functionality is helpful when, for example, return value of a
	// function should be saved and checked.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
