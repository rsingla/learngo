package calculator

import (
	"log"
	"sync"
	"time"

	"github.com/rsingla/learngo/pkg/model"
)

func AggregateAmortization(amortizationResults model.AmortizationResults) map[int]model.PaymentTable {
	amortizations := amortizationResults.PaymentPlans

	paymentMap := make(map[int]model.PaymentTable)

	//var paymentTables []PaymentTable

	for _, amortization := range amortizations {
		monthlyPayments := amortization.MonthlyPayments
		for i, monthlyPayment := range monthlyPayments {
			paymentTable := paymentMap[i]
			paymentTabs := updateMap(paymentTable, monthlyPayment)
			paymentMap[i] = *paymentTabs
		}
	}

	/*for _, paymentTable := range paymentMap {
		paymentTables = append(paymentTables, paymentTable)
	}*/

	return paymentMap

}

func updateMap(paymentTable model.PaymentTable, monthlyPay model.MonthlyPayment) *model.PaymentTable {

	paymentTabs := new(model.PaymentTable)

	paymentTabs.Balance = paymentTable.Balance + monthlyPay.RemainingBalance
	paymentTabs.Month = monthlyPay.Month
	paymentTabs.TotalInterest = paymentTable.TotalInterest + monthlyPay.Interest
	paymentTabs.TotalPaidAmount = paymentTable.TotalPaidAmount + (monthlyPay.PrincipalPayment + monthlyPay.Interest)

	payments := paymentTable.Payment
	payment := new(model.Payment)
	payment.ID = monthlyPay.ID
	payment.PaymentAmount = (monthlyPay.PrincipalPayment + monthlyPay.Interest)
	payment.PrincipalPayment = monthlyPay.PrincipalPayment

	payments = append(payments, *payment)

	paymentTabs.Payment = payments

	return paymentTabs
}

func AllAmortizations(trades []model.Tradeline) model.AmortizationResults {
	output := make(chan model.AmortizationResults)
	input := make(chan model.AmortizationResult)
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

func handleResults(input chan model.AmortizationResult, output chan model.AmortizationResults, wg *sync.WaitGroup) {
	var results model.AmortizationResults
	for result := range input {
		results.PaymentPlans = append(results.PaymentPlans, result)
		wg.Done()
	}
	output <- results
}

func ConcurrentTradeline(trade model.Tradeline, output chan model.AmortizationResult) {
	start := time.Now()
	monthlyPayments := Amortization(trade)
	elapsed := time.Since(start)
	log.Printf("trade %s", trade.ID)
	log.Printf("ConcurrentTradeline Binomial took %s", elapsed)
	amortizationResult := model.AmortizationResult{
		MonthlyPayments: monthlyPayments}
	output <- amortizationResult
}

func GetAmortizations(trades []model.Tradeline) model.AmortizationResults {
	var results model.AmortizationResults
	for _, trade := range trades {
		start := time.Now()

		monthlyPayments := Amortization(trade)

		elapsed := time.Since(start)
		log.Printf("Binomial took %s", elapsed)

		amortizationResult := model.AmortizationResult{
			MonthlyPayments: monthlyPayments}
		results.PaymentPlans = append(results.PaymentPlans, amortizationResult)
	}

	return results
}
