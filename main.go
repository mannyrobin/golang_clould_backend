// Author:
// Studentnr:

// SOURCES
// handler test: 	https://elithrar.github.io/article/testing-http-handlers-go/
// API use(JSON): 	https://blog.alexellis.io/golang-json-api-client/
// Get port nr
// Golang's documentation pages

package main

import (
	"net/http"
	"os"
)


func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/projectinfo/v1/", HandlerGitURL)
	http.HandleFunc("/", HandlerWrongURL)
	http.ListenAndServe(":" + port, nil)
}
