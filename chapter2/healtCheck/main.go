package main

import (
	"io"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", HealthCheck)
	http.ListenAndServe(":8000", nil)
}
