package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Convert the timeHandler function to a HandlerFunc type
	th := http.HandlerFunc(handler.payoffHandler)
	health := http.HandlerFunc(handler.HandleFunc)
	// And add it to the ServeMux
	mux.Handle("/payoff", th)
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)

}
