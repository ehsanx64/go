package main

import (
	"fmt"
	"time"
)

func simpleUsage() {
	now := time.Now()
	year, month, day := now.Date()
	fmt.Println("** Simple Usage")
	fmt.Printf("Current date string: %s\n", now.String())
	fmt.Printf("Current date parts:  %d %d %d\n", year, month, day)
}

func formatting() {
	// Save current date
	now := time.Now()
	utc := now.UTC()

	fmt.Println()
	fmt.Println("** Formatting")

	// Get current date/time in RFC3339/ISO8601
	fmt.Println("Current date: \t\t" + now.Format(time.RFC3339))

	// Get current date/time in RFC3339/ISO8601 (UTC)
	fmt.Println("Current date (UTC): \t" + utc.Format(time.RFC3339))

	formatted := now.Format("2006-01-02")
	fmt.Printf("Now (format): \t\t%s\n", formatted)
}

func manipulation() {
	fmt.Println()
	fmt.Println("** Manipulation")

	now := time.Now()

	// Get the date for an hour ago
	fmt.Println("An hour ago: \t\t" + now.Add(-1*time.Hour).Format(time.RFC3339))

	// Get the date for a month ago
	fmt.Println("A month ago: \t\t" + now.AddDate(0, -1, 0).Format(time.RFC3339))

	// Get the date for 28 days ago
	fmt.Println("28 days ago: \t\t" + now.AddDate(0, 0, -28).Format(time.RFC3339))
}

func main() {
	fmt.Println("*** Date\\Time Examples")
	simpleUsage()
	formatting()
	manipulation()
}
