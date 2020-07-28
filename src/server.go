package main

import (
	"calculator"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func payoffHandler(w http.ResponseWriter, r *http.Request) {

	var trades []calculator.Tradeline
	err := json.NewDecoder(r.Body).Decode(&trades)

	fmt.Println(err)

	results := calculator.Calculate(trades)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/foo", handler)

	http.HandleFunc("/payoff", payoffHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
