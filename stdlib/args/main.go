package main

import (
	"fmt"
	"os"
)

// Iterate over os.Args and print values and their indices
func rangeOverArgs() {
	fmt.Println("# Range over args")
	for index, value := range os.Args {
		fmt.Println("index:", index, ", value:", value)
	}
}

func main() {
	fmt.Println("# Common operations")
	fmt.Println("# Number of args:", len(os.Args))
	fmt.Println("# Complete args:", os.Args)
	fmt.Println("# Program arg:", os.Args[0])
	fmt.Println("# Other args:", os.Args[1:])
	fmt.Println()
	rangeOverArgs()
}
