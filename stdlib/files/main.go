package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	TextFile = "sample.txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Slurp the file content at once
func slurpFile() {
	fmt.Println("# Slurping the file ...")
	textfile, err := os.ReadFile(TextFile)
	check(err)
	fmt.Println("The text file content:", string(textfile))
}

func seekAndRead() {
	fmt.Println("# Seek and read ...")
	buffer := make([]byte, 5)
	buff := make([]byte, 4)

	f, err := os.Open(TextFile)
	check(err)

	n, err := f.Read(buffer)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(buffer))

	// Seek to 6th byte of file
	pos, err := f.Seek(6, 0)
	check(err)
	n, err = f.Read(buff)
	check(err)
	fmt.Printf("%d bytes at %02d: %v\n", n, pos, string(buff[:n]))

	// Seek to 6th byte of file
	pos, err = f.Seek(10, 0)
	check(err)
	n, err = io.ReadAtLeast(f, buff, 4)
	check(err)
	fmt.Printf("%d bytes at %d: %v\n", n, pos, string(buff[:n]))

	// Seek to the beginning of file
	_, err = f.Seek(0, 0)
	check(err)
	reader := bufio.NewReader(f)
	s, err := reader.Peek(4)
	fmt.Printf("4 bytes: %s\n", s)
}

func readFile() {
	f, err := os.Open(TextFile)
	check(err)
	defer f.Close()

	buffer := make([]byte, 5)

	for {
		n, err := f.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}
	}
}

func main() {
	slurpFile()
	seekAndRead()
	readFile()
}
