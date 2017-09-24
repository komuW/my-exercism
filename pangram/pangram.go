package pangram

import (
	"strings"
)

const testVersion = 2

// IsPangram returns true if a given string is a pangram
func IsPangram(input string) bool {
	// arrange alphabet according to letter frequency
	alphabet := []rune{'e', 't', 'a', 'o', 'i', 'n', 's', 'r', 'h', 'l', 'd', 'c', 'u', 'm', 'f', 'p', 'g', 'w', 'y', 'b', 'v', 'k', 'x', 'j', 'q', 'z'}
	lowerInput := strings.ToLower(input)
	if len(lowerInput) < 26 {
		return false
	}
	for _, v := range alphabet {
		if !strings.ContainsRune(lowerInput, v) {
			return false
		}
	}
	return true
}
