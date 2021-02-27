package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
)

func main() {

	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		number, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Fprintln(w, number)
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		number, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Fprintln(w, number)
	})

	http.ListenAndServe(":8000", newMux)
}
