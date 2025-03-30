package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	basicUsage()
	skippingValues()
	advancedUsage()
	advancedUsage2()
	goingCrazy()
	iotaInPractice()
}

func basicUsage() {
	fmt.Println("* Basic Usage")

	const (
		Red int = iota
		Orange
		Yellow
	)

	fmt.Printf("Red: %v\n", Red)
	fmt.Printf("Orange: %v\n", Orange)
	fmt.Printf("Yellow: %v\n", Yellow)

	fmt.Println()
}

func skippingValues() {
	fmt.Println("* Skipping values")

	const (
		_   int = iota // Skip the first value of 0
		Foo            // Foo = 1
		Bar            // Bar = 2
		_              // Skip 3
		_              // Skip 4
		Bin            // Bin = 5
		// A comment or a blank line will not increment the iota
		Baz // Baz = 6
	)

	fmt.Printf("Foo: %v\n", Foo)
	fmt.Printf("Bar: %v\n", Bar)
	fmt.Printf("Bin: %v\n", Bin)
	fmt.Printf("Baz: %v\n", Baz)

	fmt.Println()
}

func advancedUsage() {
	fmt.Println("* Advanced Usage")

	const (
		read   = 1 << iota // 00000001 = 1
		write              // 00000010 = 2
		remove             // 00000100 = 3

		// admin will have all of the permissions
		admin = read | write | remove
	)

	fmt.Printf("read: %v\n", read)
	fmt.Printf("write: %v\n", write)
	fmt.Printf("remove: %v\n", remove)
	fmt.Printf("admin: %v\n", admin)

	fmt.Println()
}

func advancedUsage2() {
	fmt.Println("* Advanced Usage 2")

	const (
		_  = 1 << (iota * 10) // Ignore the first value
		KB                    // Decimal: 1024
		MB                    // Decimal: 1048576
		GB                    // Decimal: 1073741824
	)

	fmt.Printf("KB: %v\n", KB)
	fmt.Printf("MB: %v\n", MB)
	fmt.Printf("GB: %v\n", GB)

	fmt.Println()
}

func goingCrazy() {
	fmt.Println("* Going Crazy with Iota")

	const (
		tomato, apple int = iota + 1, iota + 2
		orange, chevy
		ford, _
	)

	fmt.Printf("tomato: %v\n", tomato)
	fmt.Printf("apple: %v\n", apple)
	fmt.Printf("orange: %v\n", orange)
	fmt.Printf("chevy: %v\n", chevy)
	fmt.Printf("ford: %v\n", ford)

	fmt.Println()
}

func iotaInPractice() {
	fmt.Println("* Iota in Practice")

	type Genre int

	type Book struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre Genre  `json:"genre"`
	}

	const (
		Adventure Genre = iota + 1
		Comic
		Crime
		Fiction
		Fantasy
		Historical
		Horror
		Magic
		Mystery
		Philosophical
		Political
	)

	data := []byte(`{"id":1,"name":"All About Go","genre":8}`)
	book := &Book{}
	if err := json.Unmarshal(data, book); err != nil {
		fmt.Println(err)
	}

	fmt.Println("%+v", book)
	if got, exp := book.Genre, Magic; got != exp {
		fmt.Printf("Unexpected genre. got: %[1]q(%[1]d), exp %[2]q(%[2]d)", got, exp)
	}

	fmt.Println()
}
