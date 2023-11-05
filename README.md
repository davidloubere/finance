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

	fmt.Println(monthlyPeriodicity.GetPeriodsNumberInAYear())
	// 12
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

	termInYears := loanMonthsDuration / 12

	monthlyPeriodicPayment := periodicpayment.Calculate(
		loanAmount, annualInterestRate, monthlyPeriodicity, termInYears,
	)

	fmt.Printf("%.2f\n", monthlyPeriodicPayment)
	// 1037.03
}
```

### Amortization schedule

Calculate quarterly amortization schedule given:
- 2000€ as the amount of loan
- 10% in annual interest rate
- payments spread over 18 months

```go
package main

import (
	"fmt"

	"github.com/davidloubere/finance"
	"github.com/davidloubere/finance/periodicity"
)

func main() {
	var loanAmount float64 = 2000
	var loanMonthsDuration float64 = 18
	var annualInterestRate float64 = 0.1

	quarterlyPeriodicity := periodicity.New(periodicity.Quarterly)

	termInYears := loanMonthsDuration / 12

	amortizationSchedule := finance.CalculateAmortizationSchedule(
		loanAmount, annualInterestRate, quarterlyPeriodicity, termInYears,
	)

	fmt.Printf("%v\n", amortizationSchedule)
	/*
		[
		  {
		    1686.9000578762484
		    50
		    313.09994212375165
		  }
		  {
		    1365.972617199403
		    42.17250144690615
		    320.9274406768455
		  }
		  {
		    1037.0219905056363
		    34.14931542998488
		    328.95062669376676
		  }
		  {
		    699.8475981445256
		    25.925549762640912
		    337.17439236111073
		  }
		  {
		    354.24384597438706
		    17.49618995361311
		    345.60375217013853
		  }
		  {
		    -4.945377440890297e-12
		    8.856096149359644
		    354.243845974392
		  }
		]

	*/
}
```