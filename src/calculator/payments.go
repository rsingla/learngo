package calculator

const daysInYear = 365
const monthlyDays = 30

func Amortization(trade Tradeline) []MonthlyPayment {

	balance := trade.Balance
	minimumPayment := trade.MinimumPayment
	interestrate := trade.InterestRate / 100
	if trade.InterestRate == 0.0 {
		interestrate = Rate / 100
	}

	var monthlyPayments []MonthlyPayment

	month := 0
	for balance > 0 {
		dailyPeriodicRate := interestrate / daysInYear
		dailyBalance := dailyPeriodicRate * balance
		interest := dailyBalance * monthlyDays
		principalPayment := 0.00
		if minimumPayment > balance {
			principalPayment = balance
		} else {
			principalPayment = minimumPayment - interest
		}
		balance = balance - principalPayment
		month = month + 1
		monthlyPayment := buildMonthlyPayment(interest, principalPayment, balance, trade, month)
		monthlyPayments = append(monthlyPayments, monthlyPayment)
	}

	return monthlyPayments
}

func buildMonthlyPayment(interest float64, principalPayment float64, balance float64, trade Tradeline, month int) MonthlyPayment {

	monthlyPayment := MonthlyPayment{
		ID:               trade.ID,
		Month:            month,
		PrincipalPayment: float64(int(principalPayment*100)) / 100,
		Interest:         float64(int(interest*100)) / 100,
		RemainingBalance: float64(int(balance*100)) / 100}

	return monthlyPayment
}
