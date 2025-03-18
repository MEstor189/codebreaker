package input

import (
	"backend/internal/randomizer"
	"fmt"
)

func GetUserInput() string {
	var input string
	//fmt.Println("Gib deinen Code Guess ein!")
	fmt.Scanln(&input)
	return input
}

func IsValidGuess(s string) bool {
	latin := randomizer.GetLatin()
	latinSet := make(map[rune]struct{})
	for _, r := range latin {
		latinSet[r] = struct{}{}
	}
	for _, r := range s {
		if _, exists := latinSet[r]; !exists {
			return false
		}
	}
	return true
}
