package amortization

import (
	"log"

	finance_futurevalue "github.com/davidloubere/finance/annuity/futurevalue"
	finance_periodicpayment "github.com/davidloubere/finance/annuity/periodicpayment"
	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

type Amortization struct {
	Principal        float64
	InterestPayment  float64
	PrincipalPayment float64
}

func New(principal float64, interestPayment float64, principalPayment float64) Amortization {
	return Amortization{
		Principal:        principal,
		InterestPayment:  interestPayment,
		PrincipalPayment: principalPayment,
	}
}

func CalculateInterestPayment(
	presentValue float64,
	interestRate float64,
	periodicity finance_periodicity.Periodicity,
) float64 {
	var periodsNumber = periodicity.GetPeriodsNumber()

	if periodsNumber == 0.0 {
		log.Fatalf("finance: periods number cannot be equal to zero")
	}

	var termInYears float64 = 1.0 / float64(periodsNumber)

	return finance_futurevalue.FutureValue(presentValue, interestRate, periodicity, termInYears) - presentValue
}

func CalculateAmortization(
	periodicPayment float64,
	interestPayment float64,
) float64 {
	return periodicPayment - interestPayment
}

func CalculateAmortizationSchedule(
	presentValue float64,
	interestRate float64,
	periodicity finance_periodicity.Periodicity,
	termInYears float64,
) []Amortization {
	var amortizationSchedule = []Amortization{}
	var periodsNumber = periodicity.GetPeriodsNumber()
	var periodicPayment = finance_periodicpayment.PeriodicPayment(presentValue, interestRate, periodicity, termInYears)

	for i := 1; i <= int(termInYears)*periodsNumber; i++ {
		var interestPayment = CalculateInterestPayment(presentValue, interestRate, periodicity)
		var principalPayment = CalculateAmortization(periodicPayment, interestPayment)

		presentValue -= principalPayment

		var amortization = New(presentValue, interestPayment, principalPayment)

		amortizationSchedule = append(amortizationSchedule, amortization)
	}

	return amortizationSchedule
}
