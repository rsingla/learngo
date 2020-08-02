package learngo

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/rsingla/learngo/pkg/handler"
)

func main() {

	http.HandleFunc("/foo", handler.healthRespone)

	http.HandleFunc("/payoff", handler.payoffHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
