package main

import (
	"fmt"

	"github.com/ehsanx64/go-tools/greeter"
)

func main() {
	greeter.Greet()

	message := greeter.Hello("Adam")
	fmt.Println(message)
}
