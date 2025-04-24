package greeter

import "fmt"

func Greet() {
	fmt.Println("Greeting!")
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
