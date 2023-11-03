package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

// Convert an array of two bytes to a single 16-bit unsigned integer
func byteArrayToUint16() {
	fmt.Println("** Converting []byte to unit16")

	data := []byte{0x0, 0x1}
	var value uint16

	buff := bytes.NewReader(data)
	err := binary.Read(buff, binary.LittleEndian, &value)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Value:", value)
}

/*
** Convert a string to byte array
 */
func stringToByteArray() {
	fmt.Println()
	fmt.Println("** Converting string to []byte")

	str := "I Love to Go"
	data := []byte(str)
	fmt.Printf("* String: %s\n", str)
	fmt.Printf("* []byte: %+v\n", data)
}

/*
** Convert a string to unicode code point array
 */
func stringToRuneArray() {
	fmt.Println()
	fmt.Println("** Converting string to []rune")

	str := "I Love to Go"
	runes := []rune(str)
	fmt.Printf("* String: %s\n", str)
	fmt.Printf("* []rune: %+v\n", runes)
}

func demoTypeConversions() {
	fmt.Println()
	fmt.Println("** Type conversions")

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("x, y, f, z =>", x, y, f, z)
}

func main() {
	fmt.Println("*** Conversions")
	byteArrayToUint16()
	stringToByteArray()
	stringToRuneArray()
	demoTypeConversions()
}
