# os and io examples

This folder contains small examples that demonstrate common uses of Go's `os` and `io` packages.

Examples included in `os_io_examples.go`:

- Directory management: `os.MkdirAll`, `os.RemoveAll`
- File operations: `os.OpenFile`, `os.ReadFile`, `os.Create`, `os.Open`
- Buffered reading: `bufio.Scanner`
- Copying data: `io.Copy`
- MultiWriter: `io.MultiWriter`
- Temporary files: `os.CreateTemp`
- Directory listing: `os.ReadDir` and `filepath.WalkDir`
- In-memory piping: `io.Pipe`
- Error handling with `os.IsNotExist` and `os.IsPermission`

Run:

```bash
cd golang_roadmap/03_std_lib/04_os_and_io
go run os_io_examples.go
```

Notes:

- For small files, `os.ReadFile` is simple and convenient.
- For large files, prefer `bufio.Scanner` or manual buffered reads to control memory usage.
- Use `io.Copy` to efficiently transfer data between readers and writers.
- Prefer `defer file.Close()` immediately after opening files to ensure cleanup.
- Use `os.CreateTemp` for temporary file needs and `defer os.Remove(tmp.Name())` to clean up.
