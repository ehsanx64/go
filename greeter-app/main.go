package main

import (
	"fmt"

	"github.com/ehsanx64/go/greeter"
)

func main() {
	greeter.Greet()

	message := greeter.Hello("Adam")
	fmt.Println(message)
}
