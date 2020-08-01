package model

// Tradeline Input Array request for getting information on the debt accounts.
type MonthlyPayment struct {
	ID               string
	Month            int
	Interest         float64
	PrincipalPayment float64
	RemainingBalance float64
}

type AmortizationResults struct {
	PaymentPlans []AmortizationResult
}

type AmortizationResult struct {
	MonthlyPayments []MonthlyPayment
}

type PaymentTable struct {
	Month           int
	TotalInterest   float64
	Balance         float64
	TotalPaidAmount float64
	Payment         []Payment
}

type Payment struct {
	ID               string
	PaymentAmount    float64
	PrincipalPayment float64
}
