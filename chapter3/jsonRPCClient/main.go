package main

import (
	"net/http"
	"net/rpc"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/json"
)

func main() { // Create a new RPC
	servers := rpc.NewServer()                           // Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json") // Register the service by creating a new JSON
	servers.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
