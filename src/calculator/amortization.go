package calculator

import (
	"sync"
)

func AggregateAmortization(amortizationResults AmortizationResults) []PaymentTable {
	amortizations := amortizationResults.PaymentPlans

	paymentMap := make(map[int]PaymentTable)

	var paymentTables []PaymentTable

	for _, amortization := range amortizations {
		monthlyPayments := amortization.MonthlyPayments
		for i, monthlyPayment := range monthlyPayments {
			paymentTable := paymentMap[i]
			paymentMap[i] = updateMap(paymentTable, monthlyPayment)
		}
	}

	for _, paymentTable := range paymentMap {
		paymentTables = append(paymentTables, paymentTable)
	}

	return paymentTables

}

func updateMap(paymentTable PaymentTable, monthlyPayment MonthlyPayment) PaymentTable {

	paymentTable.Balance += monthlyPayment.RemainingBalance
	paymentTable.Month = monthlyPayment.Month
	paymentTable.TotalInterest += monthlyPayment.Interest
	paymentTable.TotalPaidAmount += (monthlyPayment.PrincipalPayment + monthlyPayment.Interest)
	payment := Payment{
		ID:               monthlyPayment.ID,
		PaymentAmount:    (monthlyPayment.PrincipalPayment + monthlyPayment.Interest),
		PrincipalPayment: monthlyPayment.PrincipalPayment}
	paymentTable.Payment = append(paymentTable.Payment, payment)

	return paymentTable

}

func AllAmortizations(trades []Tradeline) AmortizationResults {
	output := make(chan AmortizationResults)
	input := make(chan AmortizationResult)
	var wg sync.WaitGroup
	go handleResults(input, output, &wg)
	defer close(output)
	for _, trade := range trades {
		wg.Add(1)
		go ConcurrentTradeline(trade, input)
	}

	wg.Wait()
	close(input)
	return <-output
}

func handleResults(input chan AmortizationResult, output chan AmortizationResults, wg *sync.WaitGroup) {
	var results AmortizationResults
	for result := range input {
		results.PaymentPlans = append(results.PaymentPlans, result)
		wg.Done()
	}
	output <- results
}

func ConcurrentTradeline(trade Tradeline, output chan AmortizationResult) {
	monthlyPayments := Amortization(trade)
	amortizationResult := AmortizationResult{
		MonthlyPayments: monthlyPayments}
	output <- amortizationResult
}
