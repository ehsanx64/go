package main

import "fmt"

/*
** Struct in Go are similar to C structs, typed collection of fields. They are
** useful to create records (data grouped together)
**
** Go supports embedding of structs and interfaces to express a more seamless
** composition of types. This feature is different than the '//go:embed' compiler
** directive (introduced in Go 1.16+) to embed file/folders into the binary at
** compile time.
 */

// person struct has name (string) and age (int) fields
type person struct {
	name string
	age  int
}

// This function construct a new person struct with a given parameter (name)
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// Pointer to local variables can be returned as they will survive the scope
	// of function
	return &p
}

// This the base struct which will be used as a part of a bigger structure
type base struct {
	num int
}

// It has a method to display what's in num
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container is the top-level struct that embeds base struct into itself. Structs
// are embedded by using them without any identifier within the struct
type container struct {
	base
	str string
}

func structBasics() { // {{{
	// This creates a new struct.
	// TODO: Does parameter are determined based on the position or type???
	fmt.Println(person{"Bob", 20})

	// Fields can be named during initialization
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted values will be zero
	// TODO: Based on their initial value?
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// Same as above, but idiomatic method to encapsulate new struct creation
	fmt.Println(newPerson("Jon"))

	// struct fields can be accessed using the . operator
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// . operator can also be used with struct pointers, the pointers are
	// automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don’t have to give
	// it a name. The value can have an anonymous struct type. This technique
	// is commonly used for table-driven tests.
	cat := struct {
		name   string
		isGood bool
	}{
		"Ben",
		true,
	}

	fmt.Println("Cat:", cat)
} // }}}

func structEmbedding() {
	// When creating structs with literals, we have initialize the embedding
	// explicitly, here the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "Hey",
	}

	// Dump it
	fmt.Println("co:", co)

	// We can access the base's fields directly on 'co'
	fmt.Printf("co = {num: %v, str: %v}\n", co.num, co.str)

	// Alternatively we can reference it by full path using the embedded type name
	fmt.Println("num:", co.base.num)

	// Since container struct embeds the base struct, base methods will be
	// container methods as well. So calling embedded methods directly of the main
	// struct is possible thus:
	fmt.Println("describe():", co.describe())

	// Embedding structs with methods may be used to bestow interface
	// implementations onto other structs. Here we see that a container now
	// implements the describer interface because it embeds base.
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}

func main() {
	//structBasics()
	structEmbedding()
}
