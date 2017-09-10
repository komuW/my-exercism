package hamming

import (
	"errors"
)

const testVersion = 6

// Distance calculates the hamming distance between two strings
func Distance(a, b string) (int, error) {
	lenOfstring := len(a)
	lenOfb := len(b)
	missMatches := 0

	// return early for errors
	if lenOfstring != lenOfb {
		return -1, errors.New("the two strings are not of equal length")
	}
	for i := 0; i < lenOfstring; i++ {
		if a[i] != b[i] {
			missMatches = missMatches + 1
		}
	}
	return missMatches, nil

}
