package compare

import (
	"backend/internal/code"
	"reflect"
)

type ComparisonResult struct {
	CorrectPositions []int
}

type ComparisonResultNormal struct {
	Contains  int
	Positions int
}

func CompareEasyMode(c *code.Code, userInput string) ComparisonResult {

	var result ComparisonResult
	runesCode := []rune(c.Value)
	runesUserInput := []rune(userInput)

	for i, r := range runesUserInput {
		if i < len(c.Value) && r == runesCode[i] {
			result.CorrectPositions = append(result.CorrectPositions, i+1)
		}
	}
	return result
}

func CompareNormalMode(c *code.Code, userInput string) ComparisonResultNormal {
	charCount := 0
	posCount := 0
	runesUserInput := []rune(userInput)
	runesCode := []rune(c.Value)
	charMap := make(map[rune]int)

	for _, ch := range c.Value {
		charMap[ch]++
	}

	for i, ch := range runesUserInput {
		if i < len(c.Value) && ch == runesCode[i] {
			posCount++
		}
		if count, exists := charMap[ch]; exists && count > 0 {
			charCount++
			charMap[ch]--
		}
	}

	return ComparisonResultNormal{
		Contains:  charCount,
		Positions: posCount,
	}

}

func CompareRightGuess(c *code.Code, userInput string) bool {
	rightGuess := reflect.DeepEqual(c.Value, userInput)
	return rightGuess
}

func EvaluatedGuessIsEmpty() ComparisonResultNormal {
	cr := ComparisonResultNormal{
		Contains:  0,
		Positions: 0,
	}
	return cr
}
