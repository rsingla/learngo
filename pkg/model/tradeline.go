package model

// Tradeline Input Array request for getting information on the debt accounts.
type Tradeline struct {
	ID                      string  `json:"id"`
	AccountName             string  `json:"account_name"`
	Balance                 float64 `json:"balance"`
	MinimumPayment          int     `json:"minimum_payment"`
	InterestRate            float64 `json:"interest_rate"`
	RevolvingOrInstallement string  `json:"revolving_or_installement"`
}
