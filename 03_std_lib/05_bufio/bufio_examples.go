package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("bufio examples starting...")

	// 1) bufio.Reader.ReadString
	s := "first line\nsecond line\nthird line"
	r := bufio.NewReader(strings.NewReader(s))
	line, _ := r.ReadString('\n')
	fmt.Printf("ReadString: %q\n", line)

	// 2) Peek + ReadBytes
	if b, err := r.Peek(6); err == nil {
		fmt.Printf("Peek(6): %q\n", string(b))
	}
	rb, _ := r.ReadBytes('\n')
	fmt.Printf("ReadBytes: %q\n", string(rb))

	// 3) ReadSlice (stops at delimiter, returns including delimiter)
	r2 := bufio.NewReader(strings.NewReader("alpha,beta,gamma"))
	slice, _ := r2.ReadSlice(',')
	fmt.Printf("ReadSlice up to comma (incl): %q\n", string(slice))

	// 4) Scanner default split (whitespace)
	fmt.Println("Scanner default tokens:")
	sc := bufio.NewScanner(strings.NewReader("one two three"))
	for sc.Scan() {
		fmt.Println(" token:", sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Println("scanner error:", err)
	}

	// 5) Scanner with custom comma-split function (trims spaces)
	fmt.Println("Scanner custom comma split:")
	csv := "a, b, c ,d"
	sc2 := bufio.NewScanner(strings.NewReader(csv))
	sc2.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				tok := bytes.TrimSpace(data[:i])
				return i + 1, tok, nil
			}
		}
		if atEOF {
			tok := bytes.TrimSpace(data)
			return len(data), tok, nil
		}
		return 0, nil, nil
	})
	for sc2.Scan() {
		fmt.Printf(" token:%q\n", sc2.Text())
	}
	if err := sc2.Err(); err != nil {
		log.Println("scanner error:", err)
	}

	// 6) Buffered writer: write to a temp file then flush
	tf, err := os.CreateTemp("", "bufio_example_*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tf.Name())
	bw := bufio.NewWriter(tf)
	_, _ = bw.WriteString("line1\n")
	_, _ = bw.WriteString("line2\n")
	// data buffered; not yet on disk until Flush()
	if err := bw.Flush(); err != nil {
		log.Fatal(err)
	}
	tf.Close()
	fmt.Println("Wrote to", tf.Name())

	// 7) Read back and demonstrate ReadRune / UnreadRune
	rf, err := os.Open(tf.Name())
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(rf)
	r1, size, _ := br.ReadRune()
	fmt.Printf("ReadRune: %q size=%d\n", r1, size)
	if err := br.UnreadRune(); err != nil {
		log.Println("UnreadRune error:", err)
	}
	r3, _, _ := br.ReadRune()
	fmt.Printf("ReadRune again: %q\n", r3)
	rf.Close()

	fmt.Println("bufio examples done")
}
