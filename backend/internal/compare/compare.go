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
			result.CorrectPositions = append(result.CorrectPositions, i+1)
		}
	}
	return result
}

func CompareRightGuess(c *code.Code, userInput string) bool {
	rightGuess := reflect.DeepEqual(c.Value, userInput)
	return rightGuess
}

func EvaluatedGuessIsEmpty() ComparisonResult {
	cr := ComparisonResult{
		CorrectPositions: nil,
	}
	return cr
}
