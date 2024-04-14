package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "20")
}
