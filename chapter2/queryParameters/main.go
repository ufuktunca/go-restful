package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", queryParams["category"][0])
	fmt.Fprintf(w, "ID is: %v\n", queryParams["id"][0])
}

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/articles", QueryHandler)
	srv := &http.Server{Handler: muxRouter, Addr: "127.0.0.1:8000", WriteTimeout: 15 * time.Second, ReadTimeout: 15 * time.Second}

	srv.ListenAndServe()

}
