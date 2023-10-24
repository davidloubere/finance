package periodicity

import (
	"fmt"
	"testing"
)

func TestPeriodicityDaily(t *testing.T) {
	periodicity := New(Daily)

	var expectedPeriodeNumbers int = 365
	actualPeriodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodsNumberInAYear) {
		t.Errorf("actual %d; expected %d", actualPeriodsNumberInAYear, expectedPeriodeNumbers)
	}
}
func TestPeriodicityMonthly(t *testing.T) {
	periodicity := New(Monthly)

	var expectedPeriodeNumbers int = 12
	actualPeriodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodsNumberInAYear) {
		t.Errorf("actual %d; expected %d", actualPeriodsNumberInAYear, expectedPeriodeNumbers)
	}
}
func TestPeriodicityQuarterly(t *testing.T) {
	periodicity := New(Quarterly)

	var expectedPeriodeNumbers int = 4
	actualPeriodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodsNumberInAYear) {
		t.Errorf("actual %d; expected %d", actualPeriodsNumberInAYear, expectedPeriodeNumbers)
	}
}

func TestPeriodicitySemiAnnual(t *testing.T) {
	periodicity := New(SemiAnnual)

	var expectedPeriodeNumbers int = 2
	actualPeriodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodsNumberInAYear) {
		t.Errorf("actual %d; expected %d", actualPeriodsNumberInAYear, expectedPeriodeNumbers)
	}
}

func TestPeriodicityAnnual(t *testing.T) {
	periodicity := New(Annual)

	var expectedPeriodeNumbers int = 1
	actualPeriodsNumberInAYear := periodicity.GetPeriodsNumberInAYear()

	if fmt.Sprintf("%.d", expectedPeriodeNumbers) != fmt.Sprintf("%.d", actualPeriodsNumberInAYear) {
		t.Errorf("actual %d; expected %d", actualPeriodsNumberInAYear, expectedPeriodeNumbers)
	}
}
