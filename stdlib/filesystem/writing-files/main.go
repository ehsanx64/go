package main

import (
	"fmt"
	"os"
	"time"
)

const (
	TextFile   = "sample.txt"
	BinaryFile = "gopher-plush.jpg"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFormattedTime() string {
	return time.Now().Format(time.RFC3339)
}

func usingWriteString() {
	f, err := os.Create(TextFile)
	check(err)

	defer f.Close()

	_, err = f.WriteString(getFormattedTime() + "\n")
	check(err)

	fmt.Println("Done.")
}

func main() {
	usingWriteString()
}
