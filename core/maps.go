package main

import "fmt"

/*
** Maps are similar to PHP associate array, Perl hash, Python dictionary, Lua
** tables and so on.
 */

// The error map
var Errors = map[int]string{
	1: "Error 1",
	2: "Error 2",
}

func getByKey(code int) string {
	val, ok := Errors[code]

	if ok {
		return val
	}

	return ""
}

func basicUsage() {
	fmt.Println("** Basic Usage")

	// Create an empty map such that keys would be string and values int
	m := make(map[string]int)

	// Set some values
	m["k1"] = 7
	m["k2"] = 13

	// Dump the map
	fmt.Println("map:", m)

	// Put value of k1 into new variable v1
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	// Trying to access a non-existing key returns the default value for its type
	// In this case (map[string]int) the default value for int is zero so the v2
	v2 := m["k3"]
	fmt.Println("v2:", v2)

	// Delete k2 key and its value
	delete(m, "k2")
	fmt.Println("map:", m)

	// Accessing a non-existing key on a map returns default value (zero value)
	// of its value type.
	// In this example value type is an int, so its default value is 0
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map. This can be used to disambiguate between
	// missing keys and keys with zero values like 0 or "". Here we didn’t need
	// the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize a map in a single line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)

	// Same as above but in multile-line format
	nn := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("nn:", nn)
}

func getByKeyUsage() {
	fmt.Println("** getByKey() Usage")
	fmt.Println(getByKey(1))
	fmt.Println(getByKey(2))
}

func main() {
	fmt.Println("*** Maps Example")

	basicUsage()
	fmt.Println()
	getByKeyUsage()
}
