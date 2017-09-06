// Package twofer implements a two for one functionality.
package twofer

import (
	"fmt"
)

// ShareWith implements two for one.
func ShareWith(name string) string {

	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %v, one for me.", name)

}
