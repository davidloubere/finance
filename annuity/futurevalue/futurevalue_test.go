package futurevalue

import (
	"fmt"
	"testing"

	"github.com/davidloubere/finance/periodicity"
)

func TestFutureValueSimple(t *testing.T) {
	var presentValue float64 = 5000
	var interestRate float64 = 0.08
	var termInYears float64 = 10

	var expectedFutureValue float64 = 9000.000000
	actualFutureValueSimple := CalculateSimple(presentValue, interestRate, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", actualFutureValueSimple) {
		t.Errorf("actual %f; expected %f", actualFutureValueSimple, expectedFutureValue)
	}
}

func TestFutureValueAnnualy(t *testing.T) {
	var presentValue float64 = 5000
	var interestRate float64 = 0.05
	var termInYears float64 = 8
	periodicity := periodicity.New(periodicity.Annual)

	var expectedFutureValue float64 = 7387.277219
	actualFutureValue := Calculate(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", actualFutureValue) {
		t.Errorf("actual %f; expected %f", actualFutureValue, expectedFutureValue)
	}
}

func TestFutureValueQuarterly(t *testing.T) {
	var presentValue float64 = 2400
	var interestRate float64 = 0.06
	var termInYears float64 = 15
	periodicity := periodicity.New(periodicity.Quarterly)

	var expectedFutureValue float64 = 5863.727462
	actualFutureValue := Calculate(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", actualFutureValue) {
		t.Errorf("actual %f; expected %f", actualFutureValue, expectedFutureValue)
	}
}

func TestFutureValueMonthly(t *testing.T) {
	var presentValue float64 = 2000
	var interestRate float64 = 0.1
	var termInYears float64 = 1
	periodicity := periodicity.New(periodicity.Monthly)

	var expectedFutureValue float64 = 2209.426135
	actualFutureValue := Calculate(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", actualFutureValue) {
		t.Errorf("actual %f; expected %f", actualFutureValue, expectedFutureValue)
	}
}
