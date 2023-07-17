/*
** Enumerate serial ports if there are any
 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

const (
	defaultPort     = "/dev/ttyUSB0"
	defaultFirmware = "nodemcu"
)

// NodeMCU (lua)
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

func punyforthComm() {
	buffer := make([]byte, 100)
	commands := map[string]string{
		"hello_world": "cr println: \"Hello world!!!\" \n\r",
		"simple_arith": `
			25 2 * 3 * print: "Result: " . cr
		` + " \n\r",
		"delay_demo": `
			3000 ms 
			152 load
			+led 1000 ms -led
			create: p LED c,
			p 1 pwm-init 1023 pwm-freq 1000 pwm-duty pwm-start
			1000 ms
			1000 pwm-duty
		` + " \n\r",
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
	readBytes, err := port.Write([]byte(commands["delay_demo"]))
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

// Send the given forth file in current directory to default port
func sendForthFile(filename string) {
	buffer := make([]byte, 100)
	contents, err := ioutil.ReadFile(filename + ".fs")
	if err != nil {
		panic(err)
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
	readBytes, err := port.Write([]byte(contents))
	if err != nil {
		fmt.Println("Error writing:", err)
	} else {
		fmt.Println("Wrote", readBytes, "bytes")
		os.Exit(1)
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
	var shouldSendFile bool = false
	var infile string

	if len(os.Args) > 1 {
		shouldSendFile = true
	}

	// Display argument details if there are any
	if shouldSendFile {
		infile = os.Args[1]

		fmt.Println("Program:", os.Args[0])
		fmt.Println("Input file:", infile)
		fmt.Println("Arg count:", len(os.Args))
	}

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
		if shouldSendFile {
			sendForthFile(infile)
		} else {
			switch defaultFirmware {
			case "nodemcu":
				nodemcuComm()
			default:
				punyforthComm()
			}
		}
	} else {
		fmt.Println("Default port not found. Exiting...")
	}
}
