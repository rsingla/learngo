package finance

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/rsingla/learngo/pkg/model"
)

func MinPayoff(d model.Debt) []model.Payment {
	budget := d.MonthlyBudget
	trades := d.Tradelines
	var month int = 0

	balMap := make(map[string]float64)

	for _, trade := range trades {
		balMap[trade.ID] = trade.Balance
	}

	zeroMap := make(map[string]float64)

	amortization := []model.Payment{}

	for month <= 600 {

		myDate := addMonthsToGetYear(month)
		year := myDate.Year()
		monthly_days := daysIn(month, year)
		days := daysInAfterMonth(month)

		pays := []model.MonthlyPayment{}
		monthDate := strconv.Itoa(myDate.Year()) + " " + myDate.Month().String()

		totalInterest := 0.0
		totalPrincipal := 0.0
		totalPayment := 0.0

		for _, trade := range trades {
			balance := balMap[trade.ID]

			if balance <= 0.0 {
				continue
			}

			dailyRate := trade.InterestRate / float64(days*100)
			minPayment := float64(trade.MinimumPayment)

			monthlyPay := minimumPayment(dailyRate, monthly_days, minPayment, month, budget, balance, trade.ID)

			balMap[trade.ID] = monthlyPay.RemainingBalance

			pays = append(pays, monthlyPay)

			if balMap[trade.ID] == 0 {
				zeroMap[trade.ID] = 0.00
			}

			totalInterest += monthlyPay.Interest
			totalPrincipal += monthlyPay.PrincipalPayment
			totalPayment += (monthlyPay.Interest + monthlyPay.PrincipalPayment)
		}

		payment := model.Payment{
			Month:                 monthDate,
			TotalPayment:          totalPayment,
			TotalPrincipalPayment: totalPrincipal,
			TotalInterest:         totalInterest,
			MonthlyPayments:       pays}

		amortization = append(amortization, payment)

		if len(zeroMap) == len(balMap) {
			break
		}

		month++
	}

	fmt.Println("Total Time : ", month, amortization)

	return amortization
}

//A = P * (r(1+r)^n)/ ((1+r)^n - 1)
func minimumPayment(dailyRate float64, monthly_days int, minPayment float64, month int, budget int, balance float64, id string) model.MonthlyPayment {

	interestPayment := balance * dailyRate * float64(monthly_days)
	principalPayment := minPayment - interestPayment

	if balance < minPayment {
		principalPayment = balance
	}

	balance = balance - principalPayment

	monthlyPay := model.MonthlyPayment{ID: id, Month: month, Interest: interestPayment, PrincipalPayment: principalPayment, RemainingBalance: balance}

	fmt.Println(monthlyPay)

	return monthlyPay
}

func currMonth() time.Month {
	currentTime := time.Now()
	return currentTime.Month()
}

func daysIn(monthsAfter int, year int) int {
	myDate := time.Now().AddDate(0, monthsAfter, 0)
	return time.Date(year, myDate.Month(), 0, 0, 0, 0, 0, time.UTC).Day()
}

func daysInAfterMonth(monthsAfter int) int {
	myDate := time.Now().AddDate(0, monthsAfter, 0)
	myDateInYear := time.Now().AddDate(1, monthsAfter, 0)
	diff := (myDateInYear.Sub(myDate).Hours() / 24)
	return int(diff)
}

func addMonthsToGetYear(monthsAfter int) time.Time {
	myDate := time.Now().AddDate(0, monthsAfter, 0)
	return myDate
}

func daysInYear(year int) int {
	// This is equivalent to time.daysIn(m, year).
	days := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	return days
}

//N = (-log(1- r * A / P)) / log (1 + r)
// r is yearly Rate interest, A is total balance, P is payment
func payoffTime(rate float64, trade model.Tradeline) float64 {
	balance := float64(trade.Balance)
	minPayment := float64(trade.MinimumPayment)
	yearlyRate := rate / (12 * 100)
	payoffNumerator := -1 * math.Log(1-yearlyRate*balance/minPayment)
	payoffTimeDenominator := math.Log(1 + yearlyRate)
	payoffTime := payoffNumerator / payoffTimeDenominator

	return payoffTime
}
