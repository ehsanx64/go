package main

import (
	"fmt"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func main() {
	var t = time.Now()
	pt := ptime.New(t)

	fmt.Println("## Gregorian")
	fmt.Println(t.Date())
	fmt.Println(t.Format("2006/01/02"))
	fmt.Println()
	fmt.Println("## Jalali")
	fmt.Println(pt.Date())
	fmt.Println(pt.Format("yyyy/MM/dd"))
}
