package finance

import (
	"testing"

	"github.com/rsingla/learngo/pkg/model"
)

func Test_daysInYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{"Test DaysInYear 2020", args{2020}, 366},
		{"Test DaysInYear 2021", args{2021}, 365},
		{"Test DaysInYear 2022", args{2022}, 365},
		{"Test DaysInYear 2023", args{2023}, 365},
		{"Test DaysInYear 2024", args{2024}, 366}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysInYear(tt.args.year); got != tt.want {
				t.Errorf("daysInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_daysInYearAfterMonths(t *testing.T) {
	type args struct {
		monthsAfter int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test DaysInYear 2020", args{1}, 2020},
		{"Test DaysInYear 2021", args{2}, 2020},
		{"Test DaysInYear 2022", args{3}, 2020},
		{"Test DaysInYear 2023", args{4}, 2020},
		{"Test DaysInYear 2024", args{5}, 2021},
		{"Test DaysInYear 2024", args{6}, 2021},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addMonthsToGetYear(tt.args.monthsAfter); got != tt.want {
				t.Errorf("addMonthsToGetYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_daysIn(t *testing.T) {
	type args struct {
		monthsAfter int
		year        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test First", args{1, 2020}, 31},
		{"Test Second", args{4, 2021}, 30},
		{"Test Third", args{8, 2020}, 31},
		{"Test Fourth", args{12, 2022}, 31},
		{"Test Fifth", args{16, 2021}, 30}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysIn(tt.args.monthsAfter, tt.args.year); got != tt.want {
				t.Errorf("daysIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_daysInAfterMonth(t *testing.T) {
	type args struct {
		monthsAfter int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test daysInAfterMonth 0", args{0}, 365},
		{"Test daysInAfterMonth 1", args{1}, 365},
		{"Test daysInAfterMonth 2", args{2}, 365},
		{"Test daysInAfterMonth 9", args{9}, 365},
		{"Test daysInAfterMonth 10", args{10}, 365},
		{"Test daysInAfterMonth 100", args{39}, 366},
		{"Test daysInAfterMonth 100", args{100}, 365},
		{"Test daysInAfterMonth 200", args{200}, 365},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysInAfterMonth(tt.args.monthsAfter); got != tt.want {
				t.Errorf("daysInAfterMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinPayoff(t *testing.T) {
	type args struct {
		d model.Debt
	}

	item1 := model.Tradeline{ID: "1", AccountName: "BOA", Balance: 64367, MinimumPayment: 1100, InterestRate: 10.99, RevolvingOrInstallement: "I"}
	item2 := model.Tradeline{ID: "2", AccountName: "Chase", Balance: 40000, MinimumPayment: 1500, InterestRate: 18.99, RevolvingOrInstallement: "I"}

	trades := []model.Tradeline{}
	trades = append(trades, item1)
	trades = append(trades, item2)

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test First", args{model.Debt{MonthlyBudget: 5000, Tradelines: trades}}, "API Called"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPayoff(tt.args.d); got != tt.want {
				t.Errorf("MinPayoff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinPayoffWith5Tradelines(t *testing.T) {
	type args struct {
		d model.Debt
	}

	item1 := model.Tradeline{ID: "1", AccountName: "BOA", Balance: 64367, MinimumPayment: 1100, InterestRate: 10.99, RevolvingOrInstallement: "I"}
	item2 := model.Tradeline{ID: "2", AccountName: "Chase", Balance: 40000, MinimumPayment: 1500, InterestRate: 18.99, RevolvingOrInstallement: "I"}
	item3 := model.Tradeline{ID: "3", AccountName: "Citi", Balance: 45000, MinimumPayment: 1200, InterestRate: 21.99, RevolvingOrInstallement: "I"}
	item4 := model.Tradeline{ID: "4", AccountName: "Discover", Balance: 47000, MinimumPayment: 1700, InterestRate: 16.99, RevolvingOrInstallement: "I"}
	item5 := model.Tradeline{ID: "5", AccountName: "Amex", Balance: 49000, MinimumPayment: 1900, InterestRate: 15.99, RevolvingOrInstallement: "I"}

	trades := []model.Tradeline{}
	trades = append(trades, item1)
	trades = append(trades, item2)
	trades = append(trades, item3)
	trades = append(trades, item4)
	trades = append(trades, item5)

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 5 Tradelines", args{model.Debt{MonthlyBudget: 5000, Tradelines: trades}}, "API Called"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPayoff(tt.args.d); got != tt.want {
				t.Errorf("MinPayoff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumPayment(t *testing.T) {
	type args struct {
		dailyRate    float64
		monthly_days int
		minPayment   float64
		month        int
		budget       int
		balance      float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test First", args{dailyRate: 0.000273972602739726, monthly_days: 31, minPayment: 1100.00, month: 2, budget: 1500, balance: 64367}},
		{"Test Second", args{dailyRate: 0.000273972602739726, monthly_days: 30, minPayment: 399.00, month: 2, budget: 1000, balance: 12000}},
		{"Test Third", args{dailyRate: 0.000273972602739726, monthly_days: 29, minPayment: 300.00, month: 3, budget: 1800, balance: 11121}},
		{"Test Fourth", args{dailyRate: 0.000273972602739726, monthly_days: 28, minPayment: 299.00, month: 6, budget: 1200, balance: 12100}},
		{"Test Fifth", args{dailyRate: 0.000273972602739726, monthly_days: 30, minPayment: 300.00, month: 10, budget: 1600, balance: 13000}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minimumPayment(tt.args.dailyRate, tt.args.monthly_days, tt.args.minPayment, tt.args.month, tt.args.budget, tt.args.balance)
		})
	}
}
