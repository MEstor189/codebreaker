package round

import (
	"backend/internal/compare"
	"backend/internal/input"
	"backend/internal/level"
	"backend/internal/output"
	"fmt"
	"strings"
)

type Round struct {
	Level   level.Level
	Counter int
	Running bool
}

func CreateNewRound(counter int) *Round {
	level := level.CreateNewLevel(counter)
	r := Round{
		Level:   level,
		Counter: counter,
		Running: true,
	}
	return &r
}

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

func (r *Round) RoundLoop() {
	for r.Running {
		r.UpdateCounter()
		r.Level = level.CreateNewLevel(r.Counter)
		fmt.Println("Seed: ", r.Level.Code.Seed)
		for !r.Level.Solved {
			fmt.Println("Gib einen Guess ab.")
			userInput := input.GetUserInput()
			if input.IsValidGuess(userInput) {
				r.Level.Trys = r.Level.Trys + 1
				if compare.CompareRightGuess(&r.Level.Code, userInput) {
					r.Level.Endtimer()
					r.Level.Solved = output.OutputRightGuess(r.Level.Trys, r.Level.EndTime)
				} else {
					output.OutputWrongGuess(compare.Compare(&r.Level.Code, userInput))
				}
			}
		}
		r.PlayAgain()
	}
}

func (r *Round) UpdateCounter() {
	r.Counter = r.Counter + 1
}

func (r *Round) PlayAgain() {
	fmt.Println("Möchtest du ins nächste Level? Y/N")
	userInput := input.GetUserInput()
	if strings.Compare(userInput, "Y") == 0 {
		fmt.Println(":::: Neues Level wird generiert! ::::")
		r.Running = true
	} else {
		fmt.Println("Runde endet!")
		r.Running = false
	}
}
