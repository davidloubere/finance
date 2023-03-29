package periodicity

import (
	"fmt"
	"log"
)

const Annual string = "periodicity.annual"
const SemiAnnual string = "periodicity.semi_annual"
const Quarterly string = "periodicity.quarterly"
const Monthly string = "periodicity.monthly"
const Daily string = "periodicity.daily"

type Periodicity struct {
	key           string
	periodsNumber int
}

func New(periodicityKey string) Periodicity {
	if !keyExists(periodicityKey) {
		log.Fatalf(fmt.Sprintf("finance: unsupported periodicity \"%v\"", periodicityKey))
	}

	return Periodicity{
		key:           periodicityKey,
		periodsNumber: setPeriodsNumber(periodicityKey),
	}
}

func (periodicity Periodicity) GetPeriodsNumber() int {
	return periodicity.periodsNumber
}

func setPeriodsNumber(key string) int {
	var periodsNumber int

	if key == Annual {
		periodsNumber = 1
	} else if key == SemiAnnual {
		periodsNumber = 2
	} else if key == Quarterly {
		periodsNumber = 4
	} else if key == Monthly {
		periodsNumber = 12
	} else if key == Daily {
		periodsNumber = 365
	} else {
		log.Fatalf(fmt.Sprintf("finance: unable to set period numbers from periodicity \"%v\"", key))
	}

	return periodsNumber
}

func keyExists(key string) bool {
	for _, periodicityKey := range getKeys() {
		if periodicityKey == key {
			return true
		}
	}

	return false
}

func getKeys() []string {
	return []string{
		Annual,
		SemiAnnual,
		Quarterly,
		Monthly,
		Daily,
	}
}
