package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("/home/bars/go/src/github.com/go-restApi/static"))

	http.ListenAndServe(":8000", router)

}
