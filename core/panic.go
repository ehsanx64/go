package main

import "fmt"

func main() {
	n := panicHandler()
	fmt.Println("main received", n)

	n = panicHandlerRetval()
	fmt.Println("main received", n)

	n = noPanicHandler()
	fmt.Println("main received", n)
}

func noPanicHandler() int {
	fmt.Println()
	fmt.Println("noPanicHandler()")

	defer func() {
		fmt.Println("defer func()")
	}()

	m := 1
	fmt.Println("m :=", m)
	panic("noPanicHandler: fail")
	m = 2
	return m
}

func panicHandler() int {
	fmt.Println()
	fmt.Println("panicHandler()")

	defer func() {
		fmt.Println("defer func()")
		if err := recover(); err != nil {
			fmt.Println(">>> ", err)
		}
	}()

	m := 1
	fmt.Println("m :=", m)
	panic("panicHandler: fail")
	m = 2
	return m
}

func panicHandlerRetval() (m int) {
	fmt.Println()
	fmt.Println("panicHandlerRetval()")

	defer func() {
		fmt.Println("defer func()")
		if err := recover(); err != nil {
			fmt.Println(">>> ", err)
			m = 2
		}
	}()

	m = 1
	fmt.Println("m :=", m)
	panic("panicHandlerRetval: fail")
	m = 3
	return m
}
