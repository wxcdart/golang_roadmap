//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Run a small server that exposes pprof endpoints at /debug/pprof/
	log.Println("starting pprof server on :6060")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
