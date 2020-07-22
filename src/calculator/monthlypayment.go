package calculator

// Tradeline Input Array request for getting information on the debt accounts.
type MonthlyPayment struct {
	ID               string  `json:"id"`
	Month            int     `json:"month"`
	Interest         float64 `json:"interest"`
	PrincipalPayment float64 `json:"principal_payment"`
	RemainingBalance float64 `json:"remaining_balance"`
}

type AmortizationResults struct {
	PaymentPlans []AmortizationResult `json:"payment_plan"`
}

type AmortizationResult struct {
	MonthlyPayments []MonthlyPayment `json:"payments"`
}

type PaymentTable struct {
	Month           int       `json:"month"`
	TotalInterest   float64   `json:"total_interest"`
	Balance         float64   `json:"remaining_balance"`
	TotalPaidAmount float64   `json:"total_paid_amount"`
	Payment         []Payment `json:"payments"`
}

type Payment struct {
	ID               string  `json:"id"`
	PaymentAmount    float64 `json:"payment_amount"`
	PrincipalPayment float64 `json:"principal_paid`
}
