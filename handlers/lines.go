package handlers

import "net/http"

type Lines struct{}

func (h *Lines) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(`{"lines": 42}`))
}
