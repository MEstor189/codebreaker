package output

import (
	"backend/internal/compare"
	"fmt"
)

func OutputWrongGuess(compRes compare.ComparisonResult) {
	fmt.Println("Die richtigen Positionen sind: ", compRes)

}

func OutputRightGuess() {
	fmt.Println("Du hast den Code geknackt!")
}
