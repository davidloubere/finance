package annuity

import (
	"fmt"
	"testing"

	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

func TestFutureValueSimple(t *testing.T) {
	var presentValue float64 = 5000
	var interestRate float64 = 0.08
	var termInYears float64 = 10

	var expectedFutureValue float64 = 9000.000000
	var futureValueSimple = FutureValueSimple(presentValue, interestRate, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", futureValueSimple) {
		t.Errorf("actual %f; expected %f", futureValueSimple, expectedFutureValue)
	}
}

func TestFutureValueAnnualy(t *testing.T) {
	var presentValue float64 = 5000
	var interestRate float64 = 0.05
	var termInYears float64 = 8
	var periodicity = finance_periodicity.New(finance_periodicity.Annual)

	var expectedFutureValue float64 = 7387.277219
	var actualFutureValue = FutureValue(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", actualFutureValue) {
		t.Errorf("actual %f; expected %f", actualFutureValue, expectedFutureValue)
	}
}

func TestFutureValueQuarterly(t *testing.T) {
	var presentValue float64 = 2400
	var interestRate float64 = 0.06
	var termInYears float64 = 15
	var periodicity = finance_periodicity.New(finance_periodicity.Quarterly)

	var expectedFutureValue float64 = 5863.727462
	var actualFutureValue = FutureValue(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedFutureValue) != fmt.Sprintf("%.6f", actualFutureValue) {
		t.Errorf("actual %f; expected %f", actualFutureValue, expectedFutureValue)
	}
}
