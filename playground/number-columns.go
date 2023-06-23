/* List numbers in columns */
package main

import (
	"fmt"
)

const LIMIT = 555

func main() {
	for i := 1; i < LIMIT; i++ {
		if i != 0 && i%10 == 0 {
			fmt.Printf("%03d \n", i)
		} else {
			fmt.Printf("%03d ", i)
		}
	}
}
