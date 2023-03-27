package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi from AWS using service catalog using pipelines %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
