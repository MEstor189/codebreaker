package input

import (
	"fmt"
)

func IsValidGuess(s string, latin []rune) bool {
	latinSet := make(map[string]struct{})
	fmt.Println(s)

	for _, r := range latin {
		latinSet[string(r)] = struct{}{}
	}
	for _, r := range s {
		if _, exists := latinSet[string(r)]; !exists {
			return false
		}
	}
	return true
}
