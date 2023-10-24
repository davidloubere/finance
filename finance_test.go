package finance

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/davidloubere/finance/amortization"
	"github.com/davidloubere/finance/periodicity"
)

func TestAmortizationQuarterly(t *testing.T) {
	var presentValue float64 = 2000
	var interestRate float64 = 0.1
	var termInYears float64 = 1
	periodicity := periodicity.New(periodicity.Quarterly)

	eASchedule := []amortization.Amortization{
		amortization.New(1518.36, 50.00, 481.64),
		amortization.New(1024.69, 37.96, 493.68),
		amortization.New(518.67, 25.62, 506.02),
		amortization.New(0.00, 12.97, 518.67),
	}
	aASchedule := CalculateAmortizationSchedule(
		presentValue, interestRate, periodicity, termInYears,
	)

	if false == slices.EqualFunc(eASchedule, aASchedule, func(eA amortization.Amortization, aA amortization.Amortization) bool {
		isInterestPaymentEqual := fmt.Sprintf("%.2f", math.Abs(eA.InterestPayment)) == fmt.Sprintf("%.2f", math.Abs(aA.InterestPayment))
		isPrincipalEqual := fmt.Sprintf("%.2f", math.Abs(eA.Principal)) == fmt.Sprintf("%.2f", math.Abs(aA.Principal))
		isPrincipalPaymentEqual := fmt.Sprintf("%.2f", math.Abs(eA.PrincipalPayment)) == fmt.Sprintf("%.2f", math.Abs(aA.PrincipalPayment))

		return isInterestPaymentEqual && isPrincipalEqual && isPrincipalPaymentEqual
	}) {
		t.Errorf("actual %v; expected %v", aASchedule, eASchedule)
	}
}
