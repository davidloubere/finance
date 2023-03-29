package main

import (
	"fmt"

	finance_annuity "github.com/davidloubere/finance/annuity"
	finance_periodicity "github.com/davidloubere/finance/periodicity"
)

func main() {
	var presentValue float64 = 2400
	var interestRate float64 = 0.06
	var termInYears float64 = 15
	var periodicity = finance_periodicity.New(finance_periodicity.Quarterly)

	var futureValue = finance_annuity.FutureValue(presentValue, interestRate, periodicity, termInYears)
	fmt.Println(
		fmt.Sprintf("%.2f", futureValue),
	)
}
