package compare

import (
	"backend/internal/code"
	"reflect"
)

type ComparisonResult struct {
	CorrectPositions []int
}

func Compare(c *code.Code, userInput string) ComparisonResult {

	var result ComparisonResult

	for i, r := range userInput {
		if i < len(c.Value) && r == rune(c.Value[i]) {
			result.CorrectPositions = append(result.CorrectPositions, i)
		}
	}
	return result
}

func CompareRightGuess(c *code.Code, userInput string) bool {
	rightGuess := reflect.DeepEqual(c.Value, userInput)
	return rightGuess
}
