package model

type Debt struct {
	MonthlyBudget int         `json:"monthly_budget"`
	Tradelines    []Tradeline `json:"tradelines"`
}
