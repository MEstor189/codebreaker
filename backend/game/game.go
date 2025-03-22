package game

import (
	"backend/internal/compare"
	"backend/internal/input"
	"backend/internal/level"
	"backend/internal/round"
	"encoding/json"
	"fmt"
)

type GuessResponse struct {
	Message        string                   `json:"message"`
	Solved         bool                     `json:"solved"`
	EvaluatedGuess compare.ComparisonResult `json:"evaluatedGuess"`
	Roundstate     round.Round              `json:"roundstate"`
}

type Game struct {
	CurrentRound *round.Round
}

func NewGame() *Game {
	return &Game{
		CurrentRound: nil,
	}
}

func (g *Game) StartGame() {
	if g.CurrentRound != nil {
		fmt.Println("Eine Runde läuft bereits.")
		return
	}
	g.CurrentRound = round.CreateNewRound(0)
	fmt.Println("Das Spiel hat begonnen!")
}

func (g *Game) IsClientInputValid(inputGuess string) bool {

	if input.IsValidGuess(inputGuess) {
		g.CurrentRound.Level.Trys = g.CurrentRound.Level.Trys + 1
		return true
	} else {
		return false
	}
}

func (g *Game) ClientGuess(clientInput string) GuessResponse {
	if g.IsClientInputValid(clientInput) {

		if compare.CompareRightGuess(&g.CurrentRound.Level.Code, clientInput) {
			evaluatedGuess := compare.Compare(&g.CurrentRound.Level.Code, clientInput)
			solved := true
			g.CurrentRound.Level.Endtimer()
			message := "Das war richtig"
			gr := GuessResponse{
				Message:        message,
				Solved:         solved,
				EvaluatedGuess: evaluatedGuess,
				Roundstate:     *g.CurrentRound,
			}
			return gr
		} else {
			evaluatedGuess := compare.Compare(&g.CurrentRound.Level.Code, clientInput)
			solved := false
			message := "Probier es nochmal"
			gr := GuessResponse{
				Message:        message,
				Solved:         solved,
				EvaluatedGuess: evaluatedGuess,
				Roundstate:     *g.CurrentRound,
			}
			return gr
		}
	} else {
		gr := GuessResponse{
			Message:        "invalid guess",
			Solved:         false,
			EvaluatedGuess: compare.EvaluatedGuessIsEmpty(),
			Roundstate:     *g.CurrentRound,
		}
		return gr
	}
}

func (g *Game) NextLevel(again bool) GuessResponse {

	if again {
		g.CurrentRound.Level = level.CreateNewLevel(g.CurrentRound.Counter)
		message := "neues level"
		solved := false
		evaluatedGuess := compare.EvaluatedGuessIsEmpty()

		gr := GuessResponse{
			Message:        message,
			Solved:         solved,
			EvaluatedGuess: evaluatedGuess,
			Roundstate:     *g.CurrentRound,
		}
		return gr
	} else {
		message := "kein neues level"
		solved := true
		evaluatedGuess := compare.EvaluatedGuessIsEmpty()

		gr := GuessResponse{
			Message:        message,
			Solved:         solved,
			EvaluatedGuess: evaluatedGuess,
			Roundstate:     *g.CurrentRound,
		}
		return gr
	}

}

func SerializeGuessResponse(response GuessResponse) []byte {

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Fehler beim Serialisieren der Antwort:", err)
		return nil
	}
	return responseJSON

}

/* func Game() {
	start := round.StartRound()

	if start {
		r := round.CreateNewRound(0)
		r.RoundLoop()
	}

} */
