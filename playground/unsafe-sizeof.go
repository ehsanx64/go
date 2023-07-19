package main

import (
	"fmt"
	"unsafe"
)

type IntBool struct {
	a int64
	b bool
	c int32
}

type IntBool_A struct {
	a int64
	b bool
	c int32
}

type BoolInt struct {
	a bool
	b int64
	c int32
}

func main() {
	fmt.Println("# Size Info")
	fmt.Printf("%10s\t%s\n", "Type", "Size")
	fmt.Printf("%10s\t%d\n", "IntBool", unsafe.Sizeof(IntBool{}))
	fmt.Printf("%10s\t%d\n", "IntBool_A", unsafe.Sizeof(IntBool_A{}))
	fmt.Printf("%10s\t%d\n", "BoolInt", unsafe.Sizeof(BoolInt{}))
	fmt.Println()

	// IntBool and BoolInt are not the same
	fmt.Println("# Test 1")
	if unsafe.Sizeof(IntBool{}) == unsafe.Sizeof(BoolInt{}) {
		fmt.Println("Are the same size")
	} else {
		fmt.Println("Are not the same size")
	}
	fmt.Println()

	// IntBool and IntBool_A are not the same
	fmt.Println("# Test 2")
	if unsafe.Sizeof(IntBool{}) == unsafe.Sizeof(IntBool_A{}) {
		fmt.Println("Are the same size")
	} else {
		fmt.Println("Are not the same size")
	}
}
