package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	currentTime := time.Now()               // returns the time in the local time zone.
	fmt.Println("The time is", currentTime) // The time is 2022-02-16 17:28:33.33210964 +0300 MSK m=+0.000023434
	// `m=` value in your output:
	// The general rule is that the wall clock is for telling time and the monotonic clock is for measuring time.

	fmt.Printf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second()) // 2022-2-16 17:17:3

	// Printing and Formatting Specific Dates
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime) // The time is 2021-08-15 14:30:45.0000001 +0300 MSK
	fmt.Println()

	// Customizing Date Strings Using the Format Method
	fmt.Println(theTime.Format("2006-1-2 15:4:5")) // 2021-8-15 14:30:45
	fmt.Println()
	// if you look at the layout:
	// `15` hours - show 24-hour format
	// `03` hours - show 12-hour format

	// In this layout, include a 0 prefix on the components
	fmt.Println("Layout, include a 0 prefix", theTime.Format("2006-01-02 03:04:05")) // Layout, include a 0 prefix 2021-08-15 02:30:45
	fmt.Println()

	// Using a Pre-Defined Format, example: time.RFC3339 or time.RFC3339Nano
	fmt.Println("Format RFC3339Nano ", theTime.Format(time.RFC3339Nano)) // Format RFC3339Nano  2021-08-15T14:30:45.0000001+03:00
	fmt.Println()

	// Parsing Dates and Times in Strings
	timeString := "2021-08-15 12:30:45"
	timeParse, err := time.Parse("2006-01-02 15:04:05", timeString) // or time.ParseInLocation with time zone
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("The parse time is", timeParse) // The parse time is 2021-08-15 12:30:45 +0000 UTC
	fmt.Println()

	// Comparing Two Times
	firstTime := time.Date(2021, 12, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime) // The first time is 2021-08-15 14:30:45.0000001 +0000 UTC

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime) // The second time is 2021-12-25 16:40:55.0000002 +0000 UTC

	fmt.Println("First time before second?", firstTime.Before(secondTime)) // First time before second? true
	fmt.Println("First time after second?", firstTime.After(secondTime))   // First time after second? false
	fmt.Println("First time equal second?", firstTime.Equal(secondTime))   // First time equal second? false
	fmt.Println()

	fmt.Println("Duration between first and second time is", firstTime.Sub(secondTime)) // Duration between first and second time is -242h10m10.0000001s
	fmt.Println()

	// Adding or Subtracting Times
	theTime2 := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The time is", theTime) // The time is 2021-08-15 14:30:45.0000001 +0300 MSK

	toAdd := 24 * time.Hour
	newTime := theTime2.Add(toAdd)
	fmt.Println("The new time is", newTime) // The new time is 2021-08-16 14:30:45.0000001 +0000 UTC
}
