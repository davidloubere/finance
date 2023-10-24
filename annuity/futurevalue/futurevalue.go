package futurevalue

import (
	"log"
	"math"

	"github.com/davidloubere/finance/periodicity"
)

// Calculates future value based on simple annual interest
func CalculateSimple(
	presentValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return presentValue * (1 + (interestRate * termInYears))
}

// Calculates future value based on compounded annual interest
func Calculate(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
	termInYears float64,
) float64 {
	compoundingPeriodsNumber := float64(periodicity.GetPeriodsNumberInAYear())

	if compoundingPeriodsNumber == 0.0 {
		log.Fatalf("finance: compounding periods number cannot be equal to zero")
	}

	return presentValue * math.Pow(
		1+(interestRate/compoundingPeriodsNumber),
		(compoundingPeriodsNumber*termInYears),
	)
}
