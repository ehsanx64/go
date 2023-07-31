package main

import (
	"fmt"
	"time"
)

/*
** Deferred expressions resumed based on their proximity not their
** order of declaration.
 */
func demoDeferOrder() {
	func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")

		defer func() {
			fmt.Println("2:1")
		}()

		func() {
			fmt.Println("2:2")
		}()

		defer func() {
			fmt.Println("2:3")
		}()

		func() {
			fmt.Println("2:4")
		}()

		func() {
			fmt.Println("2:5")
		}()

		defer func() {
			fmt.Println("2:6")
		}()

		func() {
			fmt.Println("2:7")
		}()
	}()

	func() {
		fmt.Println("3")
	}()
}

/*
** State change using deffered calls is done this way
 */
func deferredStateChange() {
	fmt.Println("# Deferred Expressions")

	x := 0

	dumpX := func() {
		fmt.Printf("x: %d\n", x)
	}

	defer dumpX()

	defer func() {
		x = 4
	}()

	defer dumpX()

	dumpX()
}

/*
** Deferred calls are executed in a queue. A sleeping (or blocked)
** function can block the rest of deferred function in the queue
 */
func deferredSleeps() {
	fmt.Println("# Deferred Sleeps")

	x := 0

	dumpX := func() {
		fmt.Printf("x: %d\n", x)
	}

	defer func() {
		fmt.Println("Last")
	}()

	// This deferred call and the above one will be blocked by
	// the sleepy one
	defer dumpX()

	// The Sleepy One
	defer func() {
		time.Sleep(2 * time.Second)
		x = 4
	}()

	defer dumpX()

	dumpX()

}

func deferredWierdAlpha() (input int) {
	// This is the default
	input = 3

	defer func() {
		fmt.Println("Inside")
		//input = 5
	}()

	fmt.Println("Outside")
	return
}

func deferredWierd() (input int) {
	defer func() {
		input = 5
	}()

	return 0
}

func main() {
	//demoDeferOrder()
	//deferredStateChange()
	//deferredSleeps()

	// Wierdoz
	fmt.Println("Wierd deffer:", deferredWierd())
	fmt.Println("AlphaWierd deffer:", deferredWierdAlpha())
}
