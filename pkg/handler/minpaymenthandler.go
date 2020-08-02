package handler

import (
	"calculator"
	"encoding/json"
	"fmt"
	"net/http"
)

func MinPaymentHandler(w http.ResponseWriter, r *http.Request) {

	var trades []calculator.Tradeline
	err := json.NewDecoder(r.Body).Decode(&trades)

	fmt.Println(err)

	results := calculator.Calculate(trades)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
