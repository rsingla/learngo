package calculator

type Results struct {
	PayoffResult  []PayoffResult      `json:"result"`
	Amortization  AmortizationResults `json:"amortization"`
	PaymentTables []PaymentTable      `json:"aggregated"`
}

// Payoff structure for paying off debt
type Payoff struct {
	ID            string  `json:"id"`
	TimeInMonths  float64 `json:"time_in_months"`
	TotalAmount   float64 `json:"total_amount"`
	TotalInterest float64 `json:"total_interest"`
	LastPayment   float64 `json:"last_payment"`
	Message       string  `json:"message"`
}

type PayoffResult struct {
	Payoff Payoff `json:"payoff"`
	Err    error  `json:"error"`
}
