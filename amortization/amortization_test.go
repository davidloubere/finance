package amortization

import (
	"fmt"
	"math"
	"slices"
	"testing"

	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

func TestAmortizationQuarterly(t *testing.T) {
	var presentValue float64 = 2000
	var interestRate float64 = 0.1
	var termInYears float64 = 1
	var periodicity = finance_periodicity.New(finance_periodicity.Quarterly)

	var expectedAmortizationSchedule = []Amortization{
		New(1518.36, 50.00, 481.64),
		New(1024.69, 37.96, 493.68),
		New(518.67, 25.62, 506.02),
		New(0.00, 12.97, 518.67),
	}
	var actualAmortizationSchedule = CalculateAmortizationSchedule(
		presentValue, interestRate, periodicity, termInYears,
	)

	if false == slices.EqualFunc(expectedAmortizationSchedule, actualAmortizationSchedule, func(eA Amortization, aA Amortization) bool {
		isInterestPaymentEqual := fmt.Sprintf("%.2f", math.Abs(eA.InterestPayment)) == fmt.Sprintf("%.2f", math.Abs(aA.InterestPayment))
		isPrincipalEqual := fmt.Sprintf("%.2f", math.Abs(eA.Principal)) == fmt.Sprintf("%.2f", math.Abs(aA.Principal))
		isPrincipalPaymentEqual := fmt.Sprintf("%.2f", math.Abs(eA.PrincipalPayment)) == fmt.Sprintf("%.2f", math.Abs(aA.PrincipalPayment))

		return isInterestPaymentEqual && isPrincipalEqual && isPrincipalPaymentEqual
	}) {
		t.Errorf("actual %v; expected %v", actualAmortizationSchedule, expectedAmortizationSchedule)
	}
}
