package main

import "fmt"

/*
** A variadic function is a function which its arguments are variable. A good
** example is the fmt.Println() function.
 */

// Define a function named plus which takes two ints as input and returns an
// int as output.
func plus(a int, b int) int {
	return a + b
}

// Define a function named plusPlus which takes three ints as input.
// If a group of arguments have the same type, we can omit the type for all
// except the last one. Here a, b and c all are ints.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Define a function named that returns two ints.
func vals() (int, int) {
	return 3, 7
}

// sum() function takes an arbitrary number of int and return their sum.
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// This function (intSeq) return an anonymous function. The returned function
// closes over the variable (i) to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)

	// Get the function retvals and put them in two variables using multiple
	// assigment. Here a contains 3 and b contains 7
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Since vals() returns 2 values; if didn't need one/some of them we use
	// the blank identifier (_) to ignore unwanted values
	_, c := vals()
	fmt.Println("c:", c)

	fmt.Println("sum1:", sum(1, 2))
	fmt.Println("sum2:", sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	// If we needed to (for example) feed an array to a variadic function.
	// array name must be prefixed with '...' (to unpack the array I guess)
	fmt.Println("sum3:", sum(nums...))

	nextInt := intSeq()
	fmt.Println("1st:", nextInt())
	fmt.Println("2nd:", nextInt())
	fmt.Println("3rd:", nextInt())
	fmt.Println("4th:", nextInt())

	// A new function call, resets the counter
	newInts := intSeq()
	fmt.Println("again:", newInts())
}
