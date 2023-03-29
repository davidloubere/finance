package annuity

import (
	"fmt"
	"testing"

	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

func TestFutureValue(t *testing.T) {
	var presentValue float64 = 150000
	var interestRate float64 = 0.01
	var termInYears float64 = 20
	var periodicity = finance_periodicity.New(finance_periodicity.Monthly)

	var expectedPeriodicPayment float64 = 689.841460
	var actualPeriodicPayment = PeriodicPayment(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedPeriodicPayment) != fmt.Sprintf("%.6f", actualPeriodicPayment) {
		t.Errorf("actual %f; expected %f", actualPeriodicPayment, expectedPeriodicPayment)
	}
}
