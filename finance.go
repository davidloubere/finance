package finance

import (
	amortization "github.com/davidloubere/finance/amortization"
	interestpayment "github.com/davidloubere/finance/amortization/interestpayment"
	futurevalue "github.com/davidloubere/finance/annuity/futurevalue"
	periodicpayment "github.com/davidloubere/finance/annuity/periodicpayment"
	presentvalue "github.com/davidloubere/finance/annuity/presentvalue"
	periodicity "github.com/davidloubere/finance/periodicity"
)

func CalculateAmortizationSchedule(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
	termInYears float64,
) []amortization.Amortization {
	amortizationSchedule := []amortization.Amortization{}
	periodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()
	periodicPayment := periodicpayment.Calculate(presentValue, interestRate, periodicity, termInYears)

	for i := 1; i <= int(termInYears)*periodsNumberInAYear; i++ {
		interestPayment := interestpayment.Calculate(presentValue, interestRate, periodicity)
		principalPayment := periodicPayment - interestPayment

		presentValue -= principalPayment

		amortization := amortization.New(presentValue, interestPayment, principalPayment)

		amortizationSchedule = append(amortizationSchedule, amortization)
	}

	return amortizationSchedule
}

func CalculateFutureValueSimple(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
	termInYears float64,
) float64 {
	return futurevalue.CalculateSimple(
		presentValue, interestRate, termInYears,
	)
}

func CalculateFutureValue(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
	termInYears float64,
) float64 {
	return futurevalue.Calculate(
		presentValue, interestRate, periodicity, termInYears,
	)
}

func CalculatePeriodicPayment(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
	termInYears float64,
) float64 {
	return periodicpayment.Calculate(presentValue, interestRate, periodicity, termInYears)
}

func CalculatePresentValue(
	futureValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return presentvalue.Calculate(
		futureValue, interestRate, termInYears,
	)
}
