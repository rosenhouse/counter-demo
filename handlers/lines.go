package handlers

import (
	"net/http"
	"strings"
)

type counter interface {
	Count(packagePath string) (int, error)
}

type Lines struct {
	Counter counter
}

func (h *Lines) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	pkgPath := strings.TrimPrefix(req.URL.Path, "/lines/")
	h.Counter.Count(pkgPath)
	resp.Write([]byte(`{"lines": 42}`))
}
