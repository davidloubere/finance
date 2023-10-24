package interestpayment

import (
	"log"

	"github.com/davidloubere/finance/annuity/futurevalue"
	"github.com/davidloubere/finance/periodicity"
)

// Calculates interest payment
func Calculate(
	presentValue float64,
	interestRate float64,
	periodicity periodicity.Periodicity,
) float64 {
	periodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()

	if periodsNumberInAYear == 0.0 {
		log.Fatalf("finance: periods number cannot be equal to zero")
	}

	var termInYears float64 = 1.0 / float64(periodsNumberInAYear)

	return futurevalue.Calculate(presentValue, interestRate, periodicity, termInYears) - presentValue
}
