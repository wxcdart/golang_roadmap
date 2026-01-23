// file_io.go: Simple File I/O example — writes a few lines to a file,
// then reads the file back and prints its contents. Demonstrates error
// handling for create/open/write/read/close.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := flag.String("file", "example.txt", "file to write/read")
	flag.Parse()

	lines := []string{
		"First line",
		"Second line",
		"Third line",
	}

	// Create (or truncate) the file for writing.
	f, err := os.Create(*filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating file:", err)
		os.Exit(1)
	}

	w := bufio.NewWriter(f)
	// `range` over a slice yields (index, value). We use the blank
	// identifier `_` to ignore the index when we only need the value.
	// Here `l` receives each string from `lines` and the index is discarded.
	for _, l := range lines {
		if _, err := fmt.Fprintln(w, l); err != nil {
			fmt.Fprintln(os.Stderr, "error writing to file:", err)
			f.Close()
			os.Exit(1)
		}
	}
	if err := w.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, "error flushing writes:", err)
		f.Close()
		os.Exit(1)
	}
	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "error closing file after write:", err)
		os.Exit(1)
	}

	// Re-open the file for reading and print its contents.
	rf, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file for read:", err)
		os.Exit(1)
	}
	defer rf.Close()

	fmt.Println("Contents of", *filename+":")
	r := bufio.NewReader(rf)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					fmt.Print(line)
				}
				break
			}
			fmt.Fprintln(os.Stderr, "error reading file:", err)
			os.Exit(1)
		}
		fmt.Print(line)
	}
}
