package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rsingla/learngo/pkg/finance"
	"github.com/rsingla/learngo/pkg/model"
)

func MinPayoffHandler(w http.ResponseWriter, r *http.Request) {

	var debt model.Debt
	err := json.NewDecoder(r.Body).Decode(&debt)

	fmt.Println(err)

	results := finance.MinPayoff(debt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
