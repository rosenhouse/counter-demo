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
	lines, err := h.Counter.Count(pkgPath)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(fmt.Sprintf(`{"error": %q}`, err)))
		return
	}
	resp.Write([]byte(fmt.Sprintf(`{"lines": %d}`, lines)))
}
