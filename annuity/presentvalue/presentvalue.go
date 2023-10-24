package presentvalue

import (
	"math"
)

// Calculates present value
func Calculate(
	futureValue float64,
	interestRate float64,
	termInYears float64,
) float64 {
	return futureValue / math.Pow(1+interestRate, termInYears)
}
