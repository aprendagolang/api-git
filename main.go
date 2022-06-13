package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/healthz", health)

	http.ListenAndServe(":8080", nil)
}

func helloWorld(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(rw, "Hello %s", name)
}

func health(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "OK")
}
