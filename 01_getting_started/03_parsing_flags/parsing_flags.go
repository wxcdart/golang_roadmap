// parsing_flags.go
// Usage: go run parsing_flags.go -file=output.txt -n=5 -shout=true

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define flags
	filename := flag.String("file", "output.txt", "Output file name")
	n := flag.Int("n", 5, "Number of lines to write")
	shout := flag.Bool("shout", false, "Print lines in uppercase")
	flag.Parse()

	// Write N lines to file
	f, err := os.Create(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	for i := 1; i <= *n; i++ {
		line := fmt.Sprintf("Line %d", i)
		if *shout {
			line = strings.ToUpper(line)
		}
		_, err := fmt.Fprintln(f, line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			f.Close()
			os.Exit(1)
		}
	}
	f.Close()

	// Read and print lines from file using bufio.Scanner

	// Why not use fmt.Fscanln?
	// The previous version used fmt.Fscanln(rf, &line) in a loop, but this only reads the first word of each line (up to the first space),
	// not the entire line. Also, Fscanln returns an error (including EOF) that causes the loop to break immediately, so it may not read all lines.
	// bufio.Scanner reads each entire line and handles EOF gracefully, stopping only when all lines are read.

	rf, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer rf.Close()

	fmt.Println("Contents of", *filename)

	// BAD CODE EXAMPLE: Only prints first word of each line
	// var line string
	// for {
	//     _, err := fmt.Fscanln(rf, &line)
	//     if err != nil {
	//         break
	//     }
	//     fmt.Println(line)
	// }

	// GOOD CODE: Prints each full line
	scanner := bufio.NewScanner(rf)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each line as it appears in the file
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	}
}
