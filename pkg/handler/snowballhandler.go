package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rsingla/learngo/pkg/calculator"
	"github.com/rsingla/learngo/pkg/model"
)

func SnowballHandler(w http.ResponseWriter, r *http.Request) {

	var trades []model.Tradeline
	err := json.NewDecoder(r.Body).Decode(&trades)

	fmt.Println(err)

	results := calculator.Calculate(trades)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
