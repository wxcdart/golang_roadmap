package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

// Demonstrates file I/O and JSON handling in Go:
// - Use os.ReadFile for small files
// - Use bufio.Scanner for buffered, line-by-line reading of large files
// - Always check errors (os.IsNotExist, os.IsPermission)
// - Use encoding/json for marshal/unmarshal and streaming with Decoder/Encoder
// - Struct tags control JSON field names and options

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}

func writeSampleFile(path string) error {
    data := "first line\nsecond line\nthird line\n"
    return os.WriteFile(path, []byte(data), 0644)
}

func readSmallFile(path string) {
    fmt.Println("-- Read small file with os.ReadFile --")
    b, err := os.ReadFile(path)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("file does not exist:", path)
            return
        }
        if os.IsPermission(err) {
            fmt.Println("permission denied reading file:", path)
            return
        }
        fmt.Println("read error:", err)
        return
    }
    fmt.Printf("contents of %s:\n%s\n", path, string(b))
}

func scanFileByLine(path string) {
    fmt.Println("-- Buffered reading with bufio.Scanner (line-by-line) --")
    f, err := os.Open(path)
    if err != nil {
        fmt.Println("open error:", err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println("line:", line)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("scanner error:", err)
    }
}

func writeAndReadJSON(path string) {
    fmt.Println("-- JSON marshal/unmarshal and streaming --")
    people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob"}}

    // Marshal to bytes (small payload)
    b, err := json.MarshalIndent(people, "", "  ")
    if err != nil {
        fmt.Println("json marshal error:", err)
        return
    }
    if err := os.WriteFile(path, b, 0644); err != nil {
        fmt.Println("write json file error:", err)
        return
    }
    fmt.Printf("wrote JSON to %s\n", path)

    // Read and unmarshal
    rb, err := os.ReadFile(path)
    if err != nil {
        fmt.Println("read json file error:", err)
        return
    }
    var got []Person
    if err := json.Unmarshal(rb, &got); err != nil {
        fmt.Println("json unmarshal error:", err)
        return
    }
    fmt.Println("unmarshaled objects:", got)

    // Streaming decode with json.Decoder (useful for large JSON streams)
    f, err := os.Open(path)
    if err != nil {
        fmt.Println("open json for decode error:", err)
        return
    }
    defer f.Close()
    dec := json.NewDecoder(f)
    var streamed []Person
    if err := dec.Decode(&streamed); err != nil {
        fmt.Println("json decode error:", err)
        return
    }
    fmt.Println("stream-decoded objects:", streamed)

    // Encode to stdout with json.Encoder (streaming encoder)
    fmt.Println("-- json.Encoder -> stdout --")
    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "  ")
    if err := enc.Encode(people); err != nil {
        fmt.Println("json encode error:", err)
    }
}

func main() {
    // prepare paths inside this example folder
    smallPath := "sample.txt"
    jsonPath := "people.json"

    // Write a small sample file and read it with os.ReadFile
    if err := writeSampleFile(smallPath); err != nil {
        fmt.Println("failed to write sample file:", err)
        return
    }
    readSmallFile(smallPath)

    // Buffered scanning (line-by-line)
    scanFileByLine(smallPath)

    // Check handling of a missing file
    fmt.Println("-- Attempt to read missing file (error handling) --")
    readSmallFile("does_not_exist.txt")

    // JSON examples
    writeAndReadJSON(jsonPath)

    fmt.Println("\nTip: Use os.ReadFile for small files, bufio.Scanner for line-by-line reading of large files, always check errors, and use encoding/json with struct tags for API-friendly JSON.")
}
