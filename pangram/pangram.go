package pangram

import (
	"strings"
)

const testVersion = 2

// IsPangram returns true if a given string is a pangram
func IsPangram(input string) bool {
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
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
