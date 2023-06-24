package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// By convention functions return the error as the last return value and have
// the type error, a built-in interface.
func f1(arg int) (int, error) {
	// In this case, the value 42 is an error condition
	if arg == 42 {
		// If error condition is occured, return -1 as return value and return
		// an errors object as 2nd. Here we initialize an errors object
		return -1, errors.New("Can't work with 42")
	}

	// On success, retval as first and nil (error code) returned as return values
	return arg + 3, nil
}

// It's possible to use custom types as errors by implementing the Error() method
// on them.
type argError struct {
	arg  int
	prob string
}

// Implement the Error() method for our argError type
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// Build a new struct, and supplying values for fields (arg and prob)
		return -1, &argError{arg, "Can't work with it"}
	}

	// Return the non-error retvals
	return arg + 3, nil
}

// Clear the screen
func cls() {
	// This is how to execute system commands
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Display a simple header
func header(msg string) {
	fmt.Printf("*** %s\n", msg)
}

func main() {
	cls()
	header("Demonstrating errors handling using f1 method (default error handling)")
	// Iterate over our test values (7 and 42). _ is the index
	for _, i := range []int{7, 42} {
		// Execute f1 method and get its retval and at last check the error code
		if r, e := f1(i); e != nil {
			// An error code value of nil denotes error, so display relevant message
			fmt.Println("f1 failed:", e)
		} else {
			// A non-nil error code is success, so we're just fine
			fmt.Println("f1 worked:", r)
		}
	}

	header("Demonstrating errors handling using f2 method (custom error handling)")
	// Iterate over test values
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to programmatically use the data in a custom error, you’ll need to get
	// the error as an instance of the custom error type via type assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
