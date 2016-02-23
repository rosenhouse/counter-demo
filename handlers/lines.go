package handlers

import (
	"fmt"
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
	lines, _ := h.Counter.Count(pkgPath)
	resp.Write([]byte(fmt.Sprintf(`{"lines": %d}`, lines)))
}
