package raindrops

import "strconv"

const testVersion = 3

// Convert transforms an interger into Rainspeak
func Convert(x int) string {
	var result string
	if x%3 == 0 {
		result = result + "Pling"
	}
	if x%5 == 0 {
		result = result + "Plang"
	}
	if x%7 == 0 {
		result = result + "Plong"
	}

	if result == "" {
		return strconv.Itoa(x)
	}

	return result

}
