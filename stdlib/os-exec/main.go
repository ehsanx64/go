package main

/*
** os/exec package can be used to execute system (OS) commands, supply stdin to
** and capture the output
 */

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

const forthCode = `.( Result: ) 25 dup + 2 * . bye`

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Simply run a system command
func runMousepad() {
	cmd := exec.Command("mousepad")
	err := cmd.Run()

	check(err)
}

// Pass a string (stdin) to a command and capture its output
func runShellCommands() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	check(err)

	fmt.Printf("translated phrase: %q\n", out.String())
}

// Feed forth code to SwiftForth
func runSfCode() {
	var out bytes.Buffer

	cmd := exec.Command("sf")
	cmd.Stdin = strings.NewReader(forthCode)
	cmd.Stdout = &out

	err := cmd.Run()
	check(err)

	fmt.Print(out.String())
}

// Feed some Forth code to gforth
func runGforthCode() {
	var out bytes.Buffer

	cmd := exec.Command("gforth", "-e", forthCode)
	cmd.Stdout = &out

	err := cmd.Run()
	check(err)

	fmt.Println(out.String())
}

// Run a command with multiple arguments
func runWithMultipleArgs() {
	program := "echo"
	arg1 := "Hello"
	arg2 := "there"
	arg3 := "folks"
	arg4 := "!!!"

	cmd := exec.Command(program, arg1, arg2, arg3, arg4)

	// The Output() runs the command and returns its output
	stdout, err := cmd.Output()
	check(err)

	fmt.Printf(string(stdout))
}

// Use a pipe for stdin
func runWithStdinPipe() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	check(err)

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "Hello ")
		io.WriteString(stdin, "world")
		io.WriteString(stdin, "!!!")
	}()

	out, err := cmd.CombinedOutput()
	check(err)

	fmt.Println(string(out))
}

// Pipe output of a system command to a Go function
func runWithStdoutPipe() {
	upper := func(str string) string {
		return strings.ToUpper(str)
	}

	cmd := exec.Command("echo", "Hello world")
	stdout, err := cmd.StdoutPipe()
	check(err)

	if err := cmd.Start(); err != nil {
		check(err)
	}

	data, err := ioutil.ReadAll(stdout)
	check(err)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(upper(string(data)))
}

func main() {
	runMousepad()
	runShellCommands()
	runSfCode()
	runGforthCode()
	runWithMultipleArgs()
	runWithStdinPipe()
	runWithStdoutPipe()
}
