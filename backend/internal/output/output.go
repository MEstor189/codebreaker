package output

import (
	"backend/internal/compare"
	"fmt"
	"time"
)

func OutputWrongGuess(compRes compare.ComparisonResult) {
	fmt.Println("Die richtigen Positionen sind: ", compRes)

}

func OutputRightGuess(trys int, duration time.Duration) bool {
	fmt.Printf("Du hast den Code in %s  mit %d Trys geknackt!\n", duration, trys)
	return true
}
