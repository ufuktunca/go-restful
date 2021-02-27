package main

import (
	"net/http"
)

func main() {
	mux := &UUID{}
	http.ListenAndServe(":8000", mux)
}
