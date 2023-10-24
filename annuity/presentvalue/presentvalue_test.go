package presentvalue

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	var futureValue float64 = 2000
	var interestRate float64 = 0.03
	var termInYears float64 = 10

	var expectedPresentValue float64 = 1488.187830
	actualPresentValue := Calculate(futureValue, interestRate, termInYears)

	if fmt.Sprintf("%.6f", expectedPresentValue) != fmt.Sprintf("%.6f", actualPresentValue) {
		t.Errorf("actual %f; expected %f", actualPresentValue, expectedPresentValue)
	}
}
