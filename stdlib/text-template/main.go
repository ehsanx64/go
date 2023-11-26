package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Gender  string
	CanHunt bool
	Skills  []string
}

const personsTpl = "persons.tmpl"

// Define a shortcut for quick panics ;)
func abort(err error) {
	if err != nil {
		panic(err)
	}
}

// Define a function so data can be moved outta code's face
func getPersons() []Person {
	return []Person{
		{
			Name:    "Adam",
			Age:     30,
			Gender:  "Male",
			CanHunt: true,
			Skills:  []string{"hunting", "foraging"},
		},
		{
			Name:    "Eve",
			Age:     30,
			Gender:  "Female",
			CanHunt: false,
			Skills:  []string{"cooking", "foraging"},
		},
		{
			Name:    "John",
			Age:     3,
			Gender:  "Male",
			CanHunt: false,
			Skills:  []string{"playing"},
		},
		{
			Name:    "Alan",
			Age:     1,
			Gender:  "Male",
			CanHunt: false,
			Skills:  []string{},
		},
	}
}

func main() {
	// Load the template
	tpl, err := template.New(personsTpl).ParseFiles(personsTpl)
	abort(err)

	// Get the persons data
	persons := getPersons()

	// Feed the data into the template engine
	err = tpl.Execute(os.Stdout, persons)
	abort(err)

	// Just for reference
	fmt.Println(persons)
}
