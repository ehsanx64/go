package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func main() {
	fmt.Println("*** Conversions")
	byteArrayToUint16()
}
