# bufio examples

This folder contains examples that demonstrate common uses of the `bufio` package:

- `bufio.Reader`: `ReadString`, `ReadBytes`, `ReadSlice`, `Peek`, `ReadRune`, `UnreadRune`.
- `bufio.Writer`: buffered writes and `Flush()`.
- `bufio.Scanner`: default split and custom split functions (e.g., comma-separated tokens).
- Practical patterns: buffered file writing, buffered reading from files and strings.

Run:

```bash
cd golang_roadmap/03_std_lib/05_bufio
go run bufio_examples.go
```
