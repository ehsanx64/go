/*
 * Sprintf example
 */

package main

import "fmt"

func main() {
	test_simple()
}

func test_simple() {
	str := "item"
	separator := "-"

	fmt.Printf("%s%s%s\n", str, separator, str)
}
