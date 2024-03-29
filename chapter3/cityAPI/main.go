package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	name string
	area uint64
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tempCity := city{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)

		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		fmt.Println(tempCity)
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.name, tempCity.area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/city", postHandler)

	http.ListenAndServe(":8000", nil)
}
