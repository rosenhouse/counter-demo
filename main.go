package main

import (
	"log"
	"net/http"
)

func main() {
	handler := &CountHandler{}
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatalf("server error: %s", err)
	}
}
