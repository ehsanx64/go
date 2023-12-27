package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	SampleFile = "sample.txt"
	TextFile   = "text.txt"
	LogFile    = "log.txt"
	DummyFile  = "dummy.txt"
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
	f, err := os.Create(SampleFile)
	check(err)

	defer f.Close()

	_, err = f.WriteString(getFormattedTime() + "\n")
	check(err)

	fmt.Println("Done.")
}

func usingWriteAt() {
	f, err := os.Create(TextFile)
	check(err)

	defer f.Close()

	initialContent := []byte("Hello ")
	content := []byte("world!\n")

	_, err2 := f.Write(initialContent)
	check(err2)

	index := int64(len(initialContent))

	_, err3 := f.WriteAt(content, index)
	check(err3)

	fmt.Println("Done.")
}

func usingWriteFile() {
	data := []byte(fmt.Sprintf("%s: %s\n", getFormattedTime(), "Item 1"))
	err := ioutil.WriteFile(LogFile, data, 0644)
	check(err)

	fmt.Println("Done.")
}

func usingFprintf() {
	f, err := os.Create(DummyFile)
	check(err)

	defer f.Close()

	const name, age = "Adam Smith", 33

	n, err := fmt.Fprintln(f, name, "is", age, "years old.")
	check(err)

	m, err := fmt.Fprintf(f, "%s\n", "That's all!")
	check(err)

	fmt.Println(n+m, "bytes written")
	fmt.Println("Done.")
}

func main() {
	usingWriteString()
	usingWriteAt()
	usingWriteFile()
	usingFprintf()
}
