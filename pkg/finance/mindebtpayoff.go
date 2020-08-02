package finance

import (
	"log"

	"github.com/rsingla/learngo/pkg/model"
)

func MinPayoff(d model.Debt) string {
	log.Println(d.MonthlyBudget)
	log.Println(d.Tradelines)

	return "API Called"
}
