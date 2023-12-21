package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "User", "Name as a string")
	age := flag.Int("age", 0, "Age as an integer")
	smart := flag.Bool("smart", true, "Smart as a bool")

	// Parse the flag and copy it into a string
	var skill string
	flag.StringVar(&skill, "skill", "none", "Skill as a string")

	// Parse flags and display 'em
	flag.Parse()
	fmt.Printf("Name: %s, Age: %d, Smart: %t, Skill: %s\n",
		*name, *age, *smart, skill)
}
