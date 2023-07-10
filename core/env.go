package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var interestingVars = []string{
	"PATH",
	"GOPATH",
	"USER",
	"SHELL",
	"HOSTNAME",
	"EDITOR",
	"PWD",
}

var colorCyan = "\033[1;36m"
var colorReset = "\033[0m"

// Check if the key exists within the map
func isInteresting(key string) bool {
	for _, v := range interestingVars {
		if v == key {
			return true
		}
	}

	return false
}

func dumpPath(path string) {
	for i, v := range strings.Split(path, ":") {
		fmt.Printf("\t%02d %s\n", i, v)
	}
}

func heading(title string) {
	fmt.Println(colorCyan, "*** "+title, colorReset)
}

// List all environment variables
func demoListVars() {
	heading("Environment Variable List")

	for i, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%02d %s = %s\n", i+1, pair[0], pair[1])
	}

	fmt.Println()
}

// List interesting environment variables
func demoListInterestingVars() {
	heading("Interesting Environment Variable List")

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if isInteresting(pair[0]) {
			if pair[0] == "PATH" {
				fmt.Printf("%s:\n", pair[0])
				dumpPath(pair[1])
			} else {
				fmt.Printf("%s = %s\n", pair[0], pair[1])
			}
		}
	}

	fmt.Println()
}

// Check for existence of
func demoLookupEnv() {
	heading("Environment Variable Lookup")

	const envVar = "username"
	// Lookup an environment variable
	// We could use os.Getenv but it would not be possible to determine if a
	// variable is not set or it is set with an empty value
	val, defined := os.LookupEnv(envVar)

	// If the env-var is set print it otherwise exit to OS with an error
	if defined {
		fmt.Println("val:", val)
	} else {
		log.Fatalf("Environment variable '%s' not found\n", envVar)
	}
}

func main() {
	demoListVars()
	demoListInterestingVars()
	demoLookupEnv()
}
