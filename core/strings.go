package main

import (
	"fmt"
	"unicode/utf8"
)

/*
** In Go strings are read-only slices of bytes. The language and standard
** library treat strings specially (as containers of text encoded in UTF-8). In
** most programming languages strings are regarded as array of characters. In Go
** concept of character is called a rune (and instead of a 8-bit (bytes) value
** is an integer that represent a unicode code point.
 */

func main() {
	const str = "สวัสดี"

	fmt.Println("String:", str)

	// Since a string is a []byte, len() can be used to get length of raw bytes
	// stored within
	fmt.Println("len():", len(str))

	// Iterating over a string can be used to retrieve individual bytes
	fmt.Print("Bytes: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()

	// To count how many runes are in a string, utf8 package is helpful.
	// Note that runtime of this function depends on the size of the string, as
	// it has to decode each UTF-8 rune sequentially. Some Tai character are
	// represented by multiple UTF-8 code points, so rune count is 6 instead of
	// the expected value 4.
	fmt.Println("Rune count:", utf8.RuneCountInString(str))

	// Range also can be used to iterate over strings
	for idx, val := range str {
		fmt.Printf("%#U => [%d]\n", val, idx)
	}

	// Iteration using utf8.DecodeRuneInString function explicitly
	fmt.Println("Using utf8.DecodeRuneInString:")
	for i, w := 0, 0; i < len(str); i += w {
		val, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U => [%d]\n", val, i)
		w = width
		examineRune(val)
	}
}

func examineRune(r rune) {
	// Values enclosed in single-quotes are rune literals so we can compare them
	// directly
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
