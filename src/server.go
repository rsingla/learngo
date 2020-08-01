package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/rsingla/learngo/pkg/handler"
)



func main() {
	h := &handler.health

	http.HandleFunc("/foo", h.healthRespone)

	http.HandleFunc("/payoff", payoffHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
