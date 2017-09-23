package pangram

import (
	"strings"
)

const testVersion = 2

// IsPangram returns true if a given string is a pangram
func IsPangram(input string) bool {
	// alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	// The above is slower
	// go test -bench .
	// ok  	github.com/komuw/my-exercism/go/pangram	1.781s
	//
	// as opposed to the following that uses a slice of alphabet arranged in letter frequency:
	//https://en.wikipedia.org/wiki/Letter_frequency
	// ok github.com/komuw/my-exercism/go/pangram	1.281s
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
