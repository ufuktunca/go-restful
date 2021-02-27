package main

import (
	"fmt"
	"net/http"
)

func main() {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))

	http.ListenAndServe(":8000", nil)

}

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request")
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after request")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing mainHandler")
	w.Write([]byte("OK"))
}
