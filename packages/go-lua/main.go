package main

import (
	"github.com/Shopify/go-lua"
)

func main() {
	// Create a new lua state
	l := lua.NewState()

	// Load the standard libraries
	lua.OpenLibraries(l)

	// Load and run a lua script
	if err := lua.DoFile(l, "hello.lua"); err != nil {
		panic(err)
	}
}
