package periodicpayment

import (
	"math"

	"github.com/davidloubere/finance/periodicity"
)

// Calculates periodic payment
func Calculate(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
	termInYears float64,
) float64 {
	numberInterestPeriodsPerYear := float64(periodicity.GetPeriodsNumberInAYear())
	periodicInterestRate := interestRate / numberInterestPeriodsPerYear
	numberInterestPeriods := numberInterestPeriodsPerYear * termInYears

	return (presentValue * periodicInterestRate) / (1 - math.Pow(1+periodicInterestRate, -numberInterestPeriods))
}
