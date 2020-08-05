package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/rsingla/learngo/pkg/handler"
)

func main() {

	mux := http.NewServeMux()

	// Convert the timeHandler function to a HandlerFunc type
	payoff := http.HandlerFunc(handler.PayoffHandler)
	minPayment := http.HandlerFunc(handler.MinPaymentHandler)
	health := http.HandlerFunc(handler.HealthRespone)
	// And add it to the ServeMux
	mux.Handle("/payoff", payoff)
	mux.Handle("/health", health)
	mux.Handle("/payoff/min", minPayment)

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)

}
