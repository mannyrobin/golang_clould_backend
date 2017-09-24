// Author:
// Studentnr:

// SOURCES
// handler test: 	https://elithrar.github.io/article/testing-http-handlers-go/
// API use(JSON): 	https://blog.alexellis.io/golang-json-api-client/
// Golang's documentation pages

package main

import (
	"net/http"
)


func main() {
	http.HandleFunc("/projectinfo/v1/", HandlerGitURL)
	http.HandleFunc("/", HandlerWrongURL)
	http.ListenAndServe("localhost:8080", nil)
}
