package round

import (
	"backend/internal/input"
	"fmt"
	"strings"
)

func StartRound() bool {
	fmt.Println("Möchtest du die Runde Starten? Y/N")
	userInput := input.GetUserInput()
	if strings.Compare(userInput, "Y") == 0 {
		fmt.Println("Runde startet => Code wird generiert!")
		return true
	} else {
		fmt.Println("Runde startet nicht!")
		return false
	}
}
