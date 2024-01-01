package main

import (
	"fmt"
	"strings"
)

// I've copied the code from cobra/cobra.go file to play with it

// ld compares two strings and returns the levenshtein distance between them.
func ld(s, t string, ignoreCase bool) int {
	if ignoreCase {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}
	d := make([][]int, len(s)+1)

	dump := func(title string) {
		fmt.Println(title)
		fmt.Println(d)
		fmt.Println()
	}

	for i := range d {
		d[i] = make([]int, len(t)+1)
		d[i][0] = i
	}
	dump("i range")

	for j := range d[0] {
		d[0][j] = j
	}
	dump("j range")

	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			dump(fmt.Sprintf("ij range, i = %d, j = %d\n", i, j))
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}

func main() {
	src := "john"
	tgt := "jack"
	res := ld(src, tgt, true)
	fmt.Println("src =>", src)
	fmt.Println("tgt =>", tgt)
	fmt.Println("lvd =>", res)
}
