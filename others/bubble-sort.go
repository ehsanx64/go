/*
** Source: https://github.com/sab-07/go_algorithms/blob/main/bubble_sort.go
 */
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Channel and WaitGroup for goroutine communication
var (
	ch = make(chan *[]int)
	wg sync.WaitGroup
)

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	orange  = "\033[38;5;208m"
	magenta = "\033[35m"
	blink   = "\033[5m"
	cyan    = "\033[36m"
)

// Generator generates 100 randow nums and append to list
func generateRandomList(lists *[]int) {
	for i := 0; i < 10; i++ {
		*lists = append(*lists, rand.Intn(1000))
	}

	fmt.Printf(yellow+"Generated:\n%v\n\n"+reset, *lists)

	// ch and lists are of same type. So no need to have * or &
	ch <- lists
	wg.Done()
}

// bubble func do the sorting
func bubbleSort() {
	lists := <-ch
	fmt.Printf(green+"Input to bubble sort:\n%v\n"+reset, *lists)
	count := 0
	lenOfList := len(*lists)

	for i := 0; i < lenOfList; i++ {
		for j := 0; j < lenOfList-1; j++ {
			count++
			perc := (float64(count)) / float64(lenOfList*(lenOfList-1)) * 100
			fmt.Printf("\033[6;1HProgress:%v", math.Floor(perc))
			fmt.Printf("\033[7;12H%v", "    ")
			fmt.Printf(orange+"\033[7;1H\033[K[%v %v]"+reset, (*lists)[j], (*lists)[j+1])

			fmt.Printf(cyan+"\033[8;1HSorting:%v", *lists)
			if (*lists)[j] > (*lists)[j+1] {
				fmt.Printf(magenta + "\r\033[7;11H--> swap " + reset)
				(*lists)[j], (*lists)[j+1] = (*lists)[j+1], (*lists)[j]
				fmt.Printf(cyan+"\033[8;1HSorting:%v", *lists)
			}
			time.Sleep(time.Millisecond * 200)
		}
	}

	wg.Done()
}

func main() {
	var lists []int

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	wg.Add(2)
	go generateRandomList(&lists)
	go bubbleSort()
	wg.Wait()

	for i := 0; i < 1; i++ {
		fmt.Printf(blink+red+"\rSorted:%v\n"+reset, lists)
	}

	time.Sleep(time.Second * 5)
}
