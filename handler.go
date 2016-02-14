package main

import (
	"fmt"
	"log"
	"net/http"
)

type CountHandler struct{}

func (h *CountHandler) ServeHTTP(
	resp http.ResponseWriter, req *http.Request) {

	packageRoot := req.URL.Path

	counter := &Counter{}
	linesOfCode, err := counter.CountLines(packageRoot)
	if err != nil {
		log.Printf("counting lines: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = resp.Write([]byte(fmt.Sprintf("%d", linesOfCode)))
	if err != nil {
		log.Printf("writing response body: %s", err)
	}
}
