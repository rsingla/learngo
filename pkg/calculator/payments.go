package calculator

import (
	"log"
	"math"
	"time"

	"github.com/rsingla/learngo/pkg/model"
)

const daysInYear = 365
const monthlyDays = 30

func Amortization(trade model.Tradeline) []model.MonthlyPayment {

	balance := trade.Balance
	minimumPayment := trade.MinimumPayment

	interestRate := trade.InterestRate
	if math.IsNaN(interestRate) || interestRate == 0.00 {
		interestRate = Rate
	}

	payoffT := payoffTime(interestRate, trade)

	if math.IsNaN(payoffT) {
		payoffT = 600
	}

	months := int(math.Ceil(payoffT))

	normalizedRate := interestRate / 100

	monthlyPayments := make([]model.MonthlyPayment, months)

	month := 0

	dailyPeriodicRate := normalizedRate / daysInYear

	start := time.Now()

	bal := float64(balance)
	minPayment := float64(minimumPayment)
	for bal > 0 {
		dailyBalance := dailyPeriodicRate * bal
		interest := dailyBalance * monthlyDays
		principalPayment := 0.00
		if minPayment > bal {
			principalPayment = bal
		} else {
			principalPayment = minPayment - interest
		}
		bal = bal - principalPayment
		monthlyPayment := buildMonthlyPayment(interest, principalPayment, bal, trade, month)
		monthlyPayments[month] = *monthlyPayment

		month = month + 1
		if month >= 599 {
			break
		}
	}

	elapsed := time.Since(start)
	log.Printf("Trade ID %s", trade.ID)
	log.Printf("Amortization Binomial took %s", elapsed)

	return monthlyPayments
}

func buildMonthlyPayment(interest float64, principalPayment float64, balance float64, trade model.Tradeline, month int) *model.MonthlyPayment {
	monthlyPayment := new(model.MonthlyPayment)
	monthlyPayment.ID = trade.ID
	monthlyPayment.Month = month
	monthlyPayment.PrincipalPayment = principalPayment
	monthlyPayment.Interest = interest
	monthlyPayment.RemainingBalance = balance

	return monthlyPayment
}
