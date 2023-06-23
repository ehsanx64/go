/*
** Enumerate serial ports if there are any
 */
package main

import (
	"fmt"
	"log"
	"os"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

const (
	defaultPort = "COM39"
)

func nodemcuComm() {
	buffer := make([]byte, 100)
	commands := map[string]string{
		"version": "print(_VERSION)\n\r",
	}

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(defaultPort, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	fmt.Println("Port", defaultPort, "is opened")

	// port.Write requires a []byte so cast the string
	readBytes, err := port.Write([]byte(commands["version"]))
	if err != nil {
		fmt.Println("Error writing:", err)
	} else {
		fmt.Println("Wrote", readBytes, "bytes")
	}

	for {
		count, err := port.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
		} else if count != 0 {
			fmt.Println(">>>", string(buffer[:count]))
		}
	}
}

func main() {
	var defaultPortFound bool = false

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}

	// Return error code to OS and exit
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		os.Exit(-1)
	}

	// Iterate over found ports
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}

		// If default port is found set the flag
		if port.Name == defaultPort {
			defaultPortFound = true
		}

	}

	if defaultPortFound {
		nodemcuComm()
	} else {
		fmt.Println("Default port not found. Exiting...")
	}
}
