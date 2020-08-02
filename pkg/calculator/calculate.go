package calculator

import (
	"errors"
	"log"
	"math"
	"time"

	"github.com/rsingla/learngo/pkg/model"
)

const Rate = 20.3

func Calculate(tradelines []model.Tradeline) model.Results {

	payoffs := make([]model.PayoffResult, len(tradelines))

	start := time.Now()

	for i, trade := range tradelines {
		interestRate := trade.InterestRate
		if math.IsNaN(interestRate) || interestRate == 0.00 {
			interestRate = Rate
		}

		payoff := tradelinePayment(interestRate, trade)

		payoffResult := model.PayoffResult{
			Payoff: payoff,
		}

		payoffs[i] = payoffResult
	}

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	start = time.Now()

	amortizations := AllAmortizations(tradelines)

	elapsed = time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	//fmt.Println(amortizations)

	start = time.Now()
	paymentTables := AggregateAmortization(amortizations)

	elapsed = time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	results := model.Results{
		PayoffResult: payoffs,
		//Amortization:  amortizations,
		PaymentTables: paymentTables,
	}

	return results
}

func tradelinePayment(rate float64, trade model.Tradeline) model.Payoff {

	payoffT := payoffTime(rate, trade)

	var payoffResponse model.Payoff
	isPaidOff := true
	if math.IsNaN(payoffT) {
		payoffT = 600.00
		isPaidOff = false
	}

	payoffResponse = buildPayoffResponse(payoffT, trade, isPaidOff)

	return payoffResponse
}

func payoffTime(rate float64, trade model.Tradeline) float64 {
	balance := trade.Balance
	minPayment := trade.MinimumPayment
	yearlyRate := rate / (12 * 100)
	payoffTime := (-1 * math.Log(1-((yearlyRate*float64(balance))/float64(minPayment)))) / (math.Log(1 + yearlyRate))

	return payoffTime
}

func buildErrorResponse() error {
	err := errors.New("Not a valid minimum payment, increase the amount")

	return err
}

func buildPayoffResponse(payoffT float64, trade model.Tradeline, isPaidOff bool) model.Payoff {

	balance := trade.Balance
	minPayment := trade.MinimumPayment
	totalAmount := payoffT * float64(minPayment)
	interestAmount := (payoffT * float64(minPayment)) - float64(balance)
	_, days := math.Modf(payoffT)
	lastPayment := days * float64(minPayment)

	payoffResponse := model.Payoff{
		ID:            trade.ID,
		TimeInMonths:  math.Ceil(payoffT),
		TotalAmount:   float64(int(totalAmount*100)) / 100,
		TotalInterest: float64(int(interestAmount*100)) / 100,
		LastPayment:   float64(int(lastPayment*100)) / 100,
		PaidOff:       isPaidOff}

	return payoffResponse
}
