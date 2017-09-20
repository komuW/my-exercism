package acronym

import (
	"strings"
)

const testVersion = 3

// Abbreviate return acronym of string
func Abbreviate(input string) string {
	var res []string
	var finalRes []string

	f := func(c rune) bool {
		if string(c) == "-" {
			return true
		} else if string(c) == " " {
			return true
		}
		return false
	}
	z := strings.FieldsFunc(input, f)

	for _, v := range z {
		res = append(res, strings.Title(v))
	}
	for _, v := range res {
		finalRes = append(finalRes, string(v[0]))
	}

	cool := strings.Join(finalRes, "")

	return cool

}
