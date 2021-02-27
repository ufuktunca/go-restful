package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	http.ListenAndServe(":8000", router)
}

func getCommandOutput(command string, arguments ...string) string {
	output, _ := exec.Command(command, arguments...).Output()

	return string(output)

}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/local/go/bin/go", "version")
	io.WriteString(w, response)
	return
}

func getFileContent(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("cat", param.ByName("name")))
	return
}
