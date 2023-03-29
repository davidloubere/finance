package annuity

import (
	"log"
	"math"

	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

func futureValueSimple(
	presentValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return presentValue * (1 + (interestRate * termInYears))
}

func futureValue(
	presentValue float64,
	interestRate float64,
	compoundingPeriodsNumber float64,
	termInYears float64,
) float64 {
	if compoundingPeriodsNumber == 0.0 {
		log.Fatalf("finance: compounding periods number cannot be equal to zero")
	}

	return presentValue * math.Pow(
		1+(interestRate/compoundingPeriodsNumber),
		(compoundingPeriodsNumber*termInYears),
	)
}

// Future value (simple annual interest)
func FutureValueSimple(
	presentValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return futureValueSimple(
		presentValue, interestRate, termInYears,
	)
}

// Future value (compounded annual interest)
func FutureValue(
	presentValue float64,
	interestRate float64,
	periodicity finance_periodicity.Periodicity,
	termInYears float64,
) float64 {
	return futureValue(
		presentValue, interestRate, float64(periodicity.GetPeriodsNumber()), termInYears,
	)
}
