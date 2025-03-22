package main

import "fmt"

func main() {
	slice1 := []string{"Adam", "Eve"}
	slice2 := []string{"Able", "Cain"}
	slice3 := append(slice1, slice2...)

	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
	fmt.Println("Slice3:", slice3)
}
