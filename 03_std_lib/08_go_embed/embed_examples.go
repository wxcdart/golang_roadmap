package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed static/index.html
var indexHTML string

//go:embed static/config.json
var configJSON []byte

//go:embed static/*
var staticFiles embed.FS

func main() {
	fmt.Println("go:embed examples starting...")

	// single file as string
	fmt.Println("--- index.html (string) ---")
	fmt.Println(indexHTML)

	// single file as []byte
	fmt.Println("--- config.json (bytes) ---")
	fmt.Println(string(configJSON))

	// list directory entries via fs.ReadDir
	fmt.Println("--- static directory entries ---")
	ents, err := fs.ReadDir(staticFiles, "static")
	if err != nil {
		log.Fatalf("ReadDir: %v", err)
	}
	for _, e := range ents {
		fmt.Printf("- %s (dir=%v)\n", e.Name(), e.IsDir())
	}

	// read file via embedded FS
	b, err := fs.ReadFile(staticFiles, "static/index.html")
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}
	fmt.Println("--- read index.html via fs.ReadFile ---")
	fmt.Println(string(b))

	fmt.Println("go:embed examples done")
}
