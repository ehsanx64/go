package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
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
	fmt.Println()
}

func readInChunks() {
	fmt.Println("# Read in chunks ...")
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

func readLineByLine() {
	var i int = 0
	fmt.Println()
	fmt.Println("# Read line by line ...")
	f, err := os.Open(TextFile)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i++
		fmt.Printf("%02d : %s\n", i, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func readByWords() {
	var i int = 0
	fmt.Println()
	fmt.Println("# Read by words ...")
	f, err := os.Open(TextFile)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		i++
		fmt.Printf("(%02d: %s) ", i, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println()
}

func readBinaryFile(readAll bool) {
	f, err := os.Open(BinaryFile)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 256)

	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		if !readAll {
			break
		}
	}

	fmt.Printf("%s\n", hex.Dump(buf))
}

func main() {
	slurpFile()
	seekAndRead()
	readInChunks()
	readLineByLine()
	readByWords()
	readBinaryFile(false)
}
