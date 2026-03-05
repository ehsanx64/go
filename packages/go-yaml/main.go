package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

type Person struct {
	Name     string
	Age      int
	Numbers  []int `yaml:"numbers,flow"`
	Children []Person
}

type Family struct {
	Name    string
	Count   uint
	Persons []Person
	Date    string
}

func main() {
	cain := Person{
		Name:    "Cain",
		Age:     -40,
		Numbers: []int{6, 6, 6},
	}

	abel := Person{
		Name:    "Abel",
		Age:     0,
		Numbers: []int{math.MaxInt, math.MaxInt8, math.MaxInt16},
	}

	adam := Person{
		Name:     "Adam",
		Age:      40,
		Numbers:  []int{1, 2, 10, 13},
		Children: []Person{cain, abel},
	}

	eve := Person{
		Name:     "Eve",
		Age:      30,
		Numbers:  []int{6, 9},
		Children: []Person{cain, abel},
	}

	now := time.Now()

	family := Family{
		Name:    "First Not Best",
		Persons: []Person{adam, eve},
		Count:   4,
		Date:    string(now.Format("2006-01-02_030405")),
	}

	bytes, err := yaml.MarshalWithOptions(family,
		// Global indentation
		// yaml.Indent(2),
		yaml.IndentSequence(true),
		yaml.UseLiteralStyleIfMultiline(false))
	if err != nil {
		log.Printf("Error marshalling data: %s", err)
	}

	fmt.Printf("Output:\n%s", string(bytes))

	err = os.WriteFile("family.yml", bytes, os.FileMode(os.ModePerm))
	if err != nil {
		log.Printf("Error writing output file: %s", err)
	}
}
