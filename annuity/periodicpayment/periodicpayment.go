package annuity

import (
	"math"

	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

func PeriodicPayment(
	presentValue float64,
	interestRate float64,
	periodicity finance_periodicity.Periodicity,
	termInYears float64,
) float64 {
	var numberInterestPeriodsPerYear = float64(periodicity.GetPeriodsNumber())
	var periodicInterestRate = interestRate / numberInterestPeriodsPerYear
	var numberInterestPeriods = numberInterestPeriodsPerYear * termInYears

	return (presentValue * periodicInterestRate) / (1 - math.Pow(1+periodicInterestRate, -numberInterestPeriods))
}
