package output

import (
	"backend/internal/compare"
	"fmt"
	"time"
)

func OutputWrongGuess(compRes compare.ComparisonResult) string {
	s := fmt.Sprint("Die richtigen Positionen sind: ", compRes)
	return s
}

func OutputRightGuess(trys int, duration time.Duration) (bool, string) {
	s := fmt.Sprintf("Du hast den Code in %s  mit %d Trys geknackt!\n", duration, trys)
	return true, s
}
