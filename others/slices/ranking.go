package main

/*
** Sort and rank a slice of structures based on the specific field in the
** struct
 */

import (
	"fmt"
	"sort"
)

type Record struct {
	Name           string
	Score1, Score2 float64
	WeightedScore  float64
	Rank           int
}

func main() {
	data := []Record{
		{"A", 90, 60, 0, 0},
		{"B", 70, 85, 0, 0},
		{"C", 80, 75, 0, 0},
	}

	for i := range data {
		data[i].WeightedScore = data[i].Score1*0.6 + data[i].Score2*0.4
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].WeightedScore > data[j].WeightedScore
	})

	for i := range data {
		data[i].Rank = i + 1
	}

	fmt.Println(data)
}
