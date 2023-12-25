package main

/*
** os/exec package can be used to execute system (OS) commands, supply stdin to
** and capture the output
 */

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const forthCode = `.( Result: ) 25 dup + 2 * . bye`

// Simply run a system command
func runMousepad() {
	cmd := exec.Command("mousepad")
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

// Pass a string (stdin) to a command and capture its output
func runShellCommands() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}

// Feed forth code to SwiftForth
func runSfCode() {
	var out bytes.Buffer

	cmd := exec.Command("sf")
	cmd.Stdin = strings.NewReader(forthCode)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out.String())
}

// Feed some Forth code to gforth
func runGforthCode() {
	var out bytes.Buffer

	cmd := exec.Command("gforth", "-e", forthCode)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}

func main() {
	runMousepad()
	runSfCode()
	runGforthCode()
}
