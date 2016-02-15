package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func echoHandler(resp http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("reading request body: %s", err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = resp.Write(bodyBytes)
	if err != nil {
		log.Printf("writing response body: %s", err)
	}
}
