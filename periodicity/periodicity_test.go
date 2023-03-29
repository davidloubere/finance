package periodicity

import (
	"fmt"
	"testing"
)

func TestPeriodicityDaily(t *testing.T) {
	var periodicity = New(Daily)

	var expectedPeriodeNumbers int = 365
	var actualPeriodNumbers = periodicity.GetPeriodsNumber()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodNumbers) {
		t.Errorf("actual %d; expected %d", actualPeriodNumbers, expectedPeriodeNumbers)
	}
}
func TestPeriodicityMonthly(t *testing.T) {
	var periodicity = New(Monthly)

	var expectedPeriodeNumbers int = 12
	var actualPeriodNumbers = periodicity.GetPeriodsNumber()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodNumbers) {
		t.Errorf("actual %d; expected %d", actualPeriodNumbers, expectedPeriodeNumbers)
	}
}
func TestPeriodicityQuarterly(t *testing.T) {
	var periodicity = New(Quarterly)

	var expectedPeriodeNumbers int = 4
	var actualPeriodNumbers = periodicity.GetPeriodsNumber()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodNumbers) {
		t.Errorf("actual %d; expected %d", actualPeriodNumbers, expectedPeriodeNumbers)
	}
}

func TestPeriodicitySemiAnnual(t *testing.T) {
	var periodicity = New(SemiAnnual)

	var expectedPeriodeNumbers int = 2
	var actualPeriodNumbers = periodicity.GetPeriodsNumber()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodNumbers) {
		t.Errorf("actual %d; expected %d", actualPeriodNumbers, expectedPeriodeNumbers)
	}
}

func TestPeriodicityAnnual(t *testing.T) {
	var periodicity = New(Annual)

	var expectedPeriodeNumbers int = 1
	var actualPeriodNumbers = periodicity.GetPeriodsNumber()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodNumbers) {
		t.Errorf("actual %d; expected %d", actualPeriodNumbers, expectedPeriodeNumbers)
	}
}
