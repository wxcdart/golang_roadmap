package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Consolidated os/io examples — single main that demonstrates common patterns.
func main() {
	fmt.Println("os/io examples starting...")

	base := "example_data"
	dir := filepath.Join(base, "subdir")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("MkdirAll: %v", err)
	}
	defer os.RemoveAll(base)

	// create a simple file
	fpath := filepath.Join(dir, "hello.txt")
	if err := os.WriteFile(fpath, []byte("Hello from os/io example\nLine two\nLine three\n"), 0644); err != nil {
		log.Fatalf("WriteFile: %v", err)
	}

	// Read whole file (for small files)
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}
	fmt.Printf("ReadFile output:\n%s\n", string(data))

	// Scanner: line-by-line
	rf, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("Open for scan: %v", err)
	}
	scanner := bufio.NewScanner(rf)
	fmt.Println("Scanner lines:")
	for scanner.Scan() {
		fmt.Println("->", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v", err)
	}
	rf.Close()

	// io.Copy: copy to stdout
	cf, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("Open for copy: %v", err)
	}
	fmt.Println("\nio.Copy output:")
	if _, err := io.Copy(os.Stdout, cf); err != nil {
		log.Printf("io.Copy: %v", err)
	}
	cf.Close()

	// MultiWriter: write to both stdout and a file
	mwf, err := os.OpenFile(filepath.Join(dir, "tee.txt"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("OpenFile multiwriter: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, mwf)
	fmt.Fprintln(mw, "This line goes to both stdout and tee.txt")
	mwf.Close()

	// Temporary file
	tmp, err := os.CreateTemp("", "osio_example_*.txt")
	if err != nil {
		log.Printf("CreateTemp: %v", err)
	} else {
		fmt.Println("Created temp file:", tmp.Name())
		_, _ = tmp.WriteString("temporary data\n")
		tmp.Close()
		defer os.Remove(tmp.Name())
	}

	// ReadDir
	entries, err := os.ReadDir(base)
	if err != nil {
		log.Printf("ReadDir: %v", err)
	} else {
		fmt.Println("ReadDir entries in example_data:")
		for _, e := range entries {
			fmt.Printf("- %s (dir=%v)\n", e.Name(), e.IsDir())
		}
	}

	// WalkDir
	fmt.Println("WalkDir visit:")
	_ = filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("walk error: %v\n", err)
			return nil
		}
		fmt.Printf("* %s (dir=%v)\n", path, d.IsDir())
		return nil
	})

	// io.Pipe
	r, w := io.Pipe()
	go func() {
		_, _ = fmt.Fprintln(w, "data through pipe line 1")
		_, _ = fmt.Fprintln(w, "data through pipe line 2")
		w.Close()
	}()
	fmt.Println("Pipe read:")
	pr := bufio.NewReader(r)
	for {
		line, err := pr.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("pipe read error: %v", err)
			break
		}
		fmt.Print(line)
	}

	// Error classification
	if _, err := os.ReadFile("non_existent_file.txt"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("non_existent_file.txt does not exist")
		} else if os.IsPermission(err) {
			fmt.Println("permission error reading file")
		} else {
			fmt.Printf("other read error: %v\n", err)
		}
	}

	fmt.Println("os/io examples done")
}
