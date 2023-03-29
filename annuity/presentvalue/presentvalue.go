package annuity

import (
	"math"
)

func presentValue(
	futureValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return futureValue / math.Pow(1+interestRate, termInYears)
}

func PresentValue(
	futureValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return presentValue(
		futureValue, interestRate, termInYears,
	)
}
