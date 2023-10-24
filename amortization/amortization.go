package amortization

type Amortization struct {
	Principal        float64
	InterestPayment  float64
	PrincipalPayment float64
}

// Returns a new Amortization
func New(principal float64, interestPayment float64, principalPayment float64) Amortization {
	return Amortization{
		Principal:        principal,
		InterestPayment:  interestPayment,
		PrincipalPayment: principalPayment,
	}
}
