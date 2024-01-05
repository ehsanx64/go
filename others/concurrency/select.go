package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func fileExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details
		return false
	}
}

func main() {
	const filename = "/tmp/file"
	c1 := make(chan string)
	c2 := make(chan string)
	msg := make(chan string)

	go func() {
		for {
			if fileExist(filename) {
				msg <- "File exists"

				f, err := os.ReadFile(filename)
				if err != nil {
					msg <- "Could not read the file"
				}

				msg <- fmt.Sprintf("Content: %s\n", f)
			} else {
				msg <- "File does not exists"
			}

			time.Sleep(time.Millisecond * 5000)
		}
	}()

	go func() {
		for {
			c1 <- "Every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 secs"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)

		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg3 := <-msg:
			fmt.Println("msg:", msg3)
		}
	}

}
