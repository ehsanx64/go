package main

import "fmt"

func main() {
	// Strings can be concatenated using + operator
	fmt.Println("\"g\" + \"o\" => " + "\"g" + "o\"")

	// Go coding style prefers expressions to be grouped thus:
	// fmt.Println("1 + 1 = ", 1+1)		// This is okay
	// fmt.Println("1 + 1 = ", 1 + 1)	// This is not
	fmt.Println("1 + 1 => ", 1+1)

	// Println add an extra space char values so following line ...
	fmt.Println("7.0 / 3.0 =>", 7.0/3.0)
	// outputs this: 7.0 / 3.0 = 2.3333333333333335

	// fmt.Println() can display boolean values
	fmt.Println("true && false =>", true && false)

	// Same extra space char as numerical values
	fmt.Println("true || false =>", true || false)
	fmt.Println("!true =>", !true)
}
