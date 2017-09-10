package hamming

import (
	"errors"
)

const testVersion = 6

// Distance calculates the hamming distance between two strings
func Distance(x, y string) (int, error) {
	a := []byte(x)
	b := []byte(y)

	lenOfstring := len(a)
	lenOfb := len(b)
	missMatches := make([]int, 0)

	// return early for errors
	if lenOfstring != lenOfb {
		return -1, errors.New("the two strings are not of equal length")
	}

	for i := 0; i < lenOfstring; i++ {
		if a[i] != b[i] {
			missMatches = append(missMatches, 1)
		}
	}
	numMissMatches := len(missMatches)

	if numMissMatches != 0 {
		return numMissMatches, nil
	}
	return 0, nil

}
