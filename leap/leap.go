//Package leap finds whether a year is leap
package leap

const testVersion = 3

// IsLeapYear finds whether a given year is leap
func IsLeapYear(year int) bool {
	div4 := year % 4
	div100 := year % 100
	div400 := year % 400

	if div4 != 0 {
		return false
	}
	if div100 == 0 && div400 != 0 {
		return false
	}

	return true
}
