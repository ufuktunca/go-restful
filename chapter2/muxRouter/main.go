package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticlesHandler)
	srv := &http.Server{Handler: muxRouter, Addr: "127.0.0.1:8000", WriteTimeout: 15 * time.Second, ReadTimeout: 15 * time.Second}

	srv.ListenAndServe()

}
