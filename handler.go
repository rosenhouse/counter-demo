package main

import (
	"fmt"
	"log"
	"net/http"
)

type counter interface {
	CountLines(importPath string) (int, error)
}

type CountHandler struct {
	Counter counter
}

func (h *CountHandler) ServeHTTP(
	resp http.ResponseWriter, req *http.Request) {

	packageRoot := req.URL.Path

	linesOfCode, err := h.Counter.CountLines(packageRoot)
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
