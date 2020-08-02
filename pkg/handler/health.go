package handler

import (
	"fmt"
	"net/http"
)

func HealthRespone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is working fine %s!", r.URL.Path[1:])
}
