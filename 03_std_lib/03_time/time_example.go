package main

import (
	"fmt"
	"time"
)

// Demonstrates common operations with the time package:
// - time.Now()
// - extracting Year, Month, Day, Hour, Minute, Second
// - Format with custom layout and predefined layouts
// - Parse a time string
// - Convert between UTC and Local
// - Sub to get duration between times
// - Add to compute a relative time

func main() {
	// current local time
	now := time.Now()
	fmt.Println("Now (local):", now)

	// access components
	fmt.Printf("Year: %d Month: %s Day: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %02d Minute: %02d Second: %02d\n", now.Hour(), now.Minute(), now.Second())

	// formatting: custom layout (Go's reference time: Mon Jan 2 15:04:05 MST 2006)
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("Custom format:", custom)

	// predefined format
	fmt.Println("RFC3339:", now.Format(time.RFC3339))
	fmt.Println("RFC1123:", now.Format(time.RFC1123))

	// parsing a time string
	timestr := "2026-01-19T12:34:56Z"
	parsed, err := time.Parse(time.RFC3339, timestr)
	if err != nil {
		fmt.Println("parse error:", err)
	} else {
		fmt.Println("Parsed (UTC):", parsed)
	}

	// convert between UTC and Local
	utc := now.UTC()
	local := utc.Local()
	fmt.Println("UTC now:", utc)
	fmt.Println("Converted back to Local:", local)

	// duration between two times
	if err == nil {
		// parsed is in UTC; compute difference from now
		dur := now.Sub(parsed)
		fmt.Printf("Duration between now and parsed: %v (Hours: %.2f)\n", dur, dur.Hours())

		// Add a duration to the parsed time
		later := parsed.Add(48 * time.Hour)
		fmt.Println("Parsed + 48 hours:", later)
	}

	// Using Add to find past/future times relative to now
	tomorrow := now.Add(24 * time.Hour)
	yesterday := now.Add(-24 * time.Hour)
	fmt.Println("Tomorrow:", tomorrow)
	fmt.Println("Yesterday:", yesterday)
}
