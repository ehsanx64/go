package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

const theFile = "persons.csv"

type Person struct {
	name string
	age  int
}

/*
** This function read the data and extract fields
 */
func createPersonList(data [][]string) []Person {
	var personList []Person

	for i, line := range data {
		// Omit header line
		if i > 0 {
			var rec Person

			for j, field := range line {
				if j == 0 {
					rec.name = field
				} else if j == 1 {
					val, e := strconv.Atoi(field)

					if e != nil {
						panic(e)
					}

					rec.age = val
				}
			}

			personList = append(personList, rec)
		}
	}

	return personList
}

func main() {
	// Open the file
	f, err := os.Open(theFile)
	if err != nil {
		log.Fatal(err)
	}

	// Close the file in the end
	defer f.Close()

	// Read values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert records to array of structs
	person := createPersonList(data)

	// Dump the list
	fmt.Printf("%v\n", person)
}
