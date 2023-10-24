package periodicpayment

import (
	"fmt"
	"testing"

	"github.com/davidloubere/finance/periodicity"
)

func TestFutureValue(t *testing.T) {
	var presentValue float64 = 150000
	var interestRate float64 = 0.01
	var termInYears float64 = 20
	periodicity := periodicity.New(periodicity.Monthly)

	var expectedPeriodicPayment float64 = 689.841460
	actualPeriodicPayment := Calculate(presentValue, interestRate, periodicity, termInYears)

	if fmt.Sprintf("%.6f", expectedPeriodicPayment) != fmt.Sprintf("%.6f", actualPeriodicPayment) {
		t.Errorf("actual %f; expected %f", actualPeriodicPayment, expectedPeriodicPayment)
	}
}
