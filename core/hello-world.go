// Every Go application should have a 'main' package
package main

// Import a library package.
// Using this way (one-line) only one package can be imported. To import
// serveral packages import the import statement should be repeated.
// Other method to import several packages use following syntax:
// import (
//		"fmt"
//		"pkg1"
//		"pkg2"
// )
import "fmt"

// Main method is the application's entry point
func main() {
	// Println display its arguments on the console and adds new-line
	// automatically.
	fmt.Println("Hello world!")
}

/*
** To run:
**		go run hello-world.go
 */
