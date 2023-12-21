package main

import (
	"fmt"
	"time"
)

func main() {
	// Save current date
	now := time.Now()
	utc := now.UTC()

	// Get current date/time in RFC3339/ISO8601
	fmt.Println("Current date: \t\t" + now.Format(time.RFC3339))

	// Get current date/time in RFC3339/ISO8601 (UTC)
	fmt.Println("Current date (UTC): \t" + utc.Format(time.RFC3339))

	// Get the date for an hour ago
	fmt.Println("An hour ago: \t\t" + utc.Add(-1*time.Hour).Format(time.RFC3339))

	// Get the date for a month ago
	fmt.Println("A month ago: \t\t" + utc.AddDate(0, -1, 0).Format(time.RFC3339))

	// Get the date for 28 days ago
	fmt.Println("28 days ago: \t\t" + utc.AddDate(0, 0, -28).Format(time.RFC3339))
}
