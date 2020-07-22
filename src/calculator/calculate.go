package calculator

import (
	"errors"
	"math"
)

const Rate = 20.3

func Calculate(tradelines []Tradeline) Results {

	var Payoffs []PayoffResult

	for _, trade := range tradelines {
		interestRate := trade.InterestRate
		if math.IsNaN(interestRate) || interestRate == 0.00 {
			interestRate = Rate
		}

		payoff, err := tradelinePayment(interestRate, trade)

		payoffResult := PayoffResult{
			Payoff: payoff,
			Err:    err,
		}

		Payoffs = append(Payoffs, payoffResult)
	}

	amortizations := AllAmortizations(tradelines)

	paymentTables := AggregateAmortization(amortizations)

	results := Results{
		PayoffResult:  Payoffs,
		PaymentTables: paymentTables,
	}

	return results
}

func tradelinePayment(rate float64, trade Tradeline) (Payoff, error) {

	balance := trade.Balance
	minPayment := trade.MinimumPayment
	yearlyRate := rate / (12 * 100)
	time := (-1 * math.Log(1-((yearlyRate*balance)/minPayment))) / (math.Log(1 + yearlyRate))
	var payoffResponse Payoff
	var err error
	if math.IsNaN(time) {
		err = buildErrorResponse()
	} else {
		payoffResponse = buildPayoffResponse(time, trade)
	}

	return payoffResponse, err
}

func buildErrorResponse() error {
	err := errors.New("Not a valid minimum payment, increase the amount")

	return err
}

func buildPayoffResponse(time float64, trade Tradeline) Payoff {

	balance := trade.Balance
	minPayment := trade.MinimumPayment
	totalAmount := time * minPayment
	interestAmount := (time * minPayment) - balance
	_, days := math.Modf(time)
	lastPayment := days * minPayment

	payoffResponse := Payoff{
		ID:            trade.ID,
		TimeInMonths:  math.Ceil(time),
		TotalAmount:   float64(int(totalAmount*100)) / 100,
		TotalInterest: float64(int(interestAmount*100)) / 100,
		LastPayment:   float64(int(lastPayment*100)) / 100}

	return payoffResponse
}
