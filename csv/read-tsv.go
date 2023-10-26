/*
** Read a .tsv file in a line-by-line fashion
 */
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const theFile = "persons.tsv"

type Person struct {
	name string
	age  int
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

	// Set tab character as the delimiter
	csvReader.Comma = '\t'

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Dump the read line
		fmt.Printf("%+v\n", rec)
	}
}
