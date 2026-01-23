# go:embed examples

This folder demonstrates embedding static files into Go binaries using the `embed` package (Go 1.16+).

Examples in `embed_examples.go`:

- Embed a single file as a string (`//go:embed file`) 
- Embed a file as bytes (`//go:embed file` into []byte)
- Embed a whole directory into `embed.FS` and read entries via `io/fs`
- Notes on using embedded assets with `http.FileServer` and rotation (not needed for embedded files)

Run:

```bash
cd golang_roadmap/03_std_lib/08_go_embed
go run embed_examples.go
```
