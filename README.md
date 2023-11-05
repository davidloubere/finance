# Finance

A Go library for finance.

[![Go Reference](https://pkg.go.dev/badge/github.com/davidloubere/finance.svg)](https://pkg.go.dev/github.com/davidloubere/finance)

## Examples

### Periodicity

Get the number of periods in a year for a monthly periodicity:

```go
package main

import (
	"fmt"

	"github.com/davidloubere/finance/periodicity"
)

func main() {
	monthlyPeriodicity := periodicity.New(periodicity.Monthly)

	fmt.Println(monthlyPeriodicity.GetPeriodsNumberInAYear()) // Outputs "12"
}
```

### Periodic payment

Calculate monthly periodic payment given:
- 10000€ as the amount of loan
- 8% in annual interest rate
- payments spread over 10 months

```go
package main

import (
	"fmt"

	"github.com/davidloubere/finance/annuity/periodicpayment"
	"github.com/davidloubere/finance/periodicity"
)

func main() {
	var loanAmount float64 = 10000
	var loanMonthsDuration float64 = 10
	var annualInterestRate float64 = 0.08

	monthlyPeriodicity := periodicity.New(periodicity.Monthly)

	termInYears := loanMonthsDuration / float64(monthlyPeriodicity.GetPeriodsNumberInAYear())

	monthlyPeriodicPayment := periodicpayment.Calculate(
		loanAmount, annualInterestRate, monthlyPeriodicity, termInYears,
	)

	fmt.Printf("%.2f\n", monthlyPeriodicPayment) // Outputs "1037.03"
}
```