package code

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpRequest() {
	fmt.Println("Hello world")

	resp, err := http.Get("https://tianpan.co/notes/120-designing-uber")

	fmt.Println(resp)
	fmt.Println(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)

	errors := ioutil.WriteFile("testdata/hello", body, 0644)
	if errors != nil {
		log.Fatal(err)
	}
}
