package main

import (
	"net/http"
	"log"
	"fmt"
	"html"
	"encoding/json"
)

type User struct { 
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Phone string  `json:"phone"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func payoffHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.GetBody)
	fmt.Println(r.Method)
	fmt.Println(r.Host)
	fmt.Println(r.Header)

	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	
	w.Header().Set("Content-Type", "application/json") 
      user := User {
                    Id: 1, 
                    Name: "John Doe", 
                    Email: "johndoe@gmail.com", 
                    Phone: "000099999"} 
      
     json.NewEncoder(w).Encode(user) 
}


func main() {
	http.HandleFunc("/foo", handler)

	http.HandleFunc("/payoff", payoffHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Println(r.GetBody)
		fmt.Println(r.Method)
		fmt.Println(r.Host)
		fmt.Println(r.Header)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
