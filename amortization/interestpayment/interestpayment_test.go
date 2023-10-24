package interestpayment

import (
	"fmt"
	"testing"

	"github.com/davidloubere/finance/periodicity"
)

func TestInterestPayment(t *testing.T) {
	var presentValue float64 = 5000
	var interestRate float64 = 0.08
	periodicity := periodicity.New(periodicity.Monthly)

	var expectedInterestPayment float64 = 33.333333
	actualInterestPayment := Calculate(presentValue, interestRate, periodicity)

	if fmt.Sprintf("%.6f", expectedInterestPayment) != fmt.Sprintf("%.6f", actualInterestPayment) {
		t.Errorf("actual %f; expected %f", actualInterestPayment, expectedInterestPayment)
	}
}
