package main

import (
	"fmt"
	"time"
)

func main() {
	i := 4
	// Unlike fmt.Println(), fmt.Print() does not add extra space
	fmt.Print("Write ", i, " as ")

	// Basic usage with numerics. Switch can have no 'default' so if the 'i'
	// value was, for example, 4, none of cases would match. In this example
	// that will cause a little annoyance. There would be no newline printed so
	// the next print would start printing from unintended position.
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Switching on time
	switch time.Now().Weekday() {
	// Multiple conditions can be checked for same outcome
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Switch without an expression, which is a form of if-else branching
	t := time.Now()
	switch {
	// Conditions can be non-constant, here an expression
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Switching on types (type switch)
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Not a bool or int, but a %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// Switching on strings
	color := "gree"
	switch color {
	case "red", "pink":
		fmt.Println("Color was red or pink")
	case "green":
		fmt.Println("Color was green")
	default:
		fmt.Println("Unsupported color:", color)
	}
}
