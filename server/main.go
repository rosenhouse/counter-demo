package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rosenhouse/counter-demo/counters"
	"github.com/rosenhouse/counter-demo/handlers"
)

func main() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalf("missing required env var GOPATH")
	}

	linesHandler := &handlers.Lines{
		Counter: &counters.PackageLinesCounter{
			GoPath:           goPath,
			DirectoryLister:  &counters.DirectoryLister{},
			FileLinesCounter: &counters.FileLinesCounter{},
		},
	}

	mux := http.NewServeMux()
	mux.Handle("/lines/", linesHandler)
	http.ListenAndServe(":8000", mux)
}
