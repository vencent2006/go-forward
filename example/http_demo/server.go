package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	// listen 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
