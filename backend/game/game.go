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
		fmt.Println("Round is running.")
		return
	}
	g.CurrentRound = round.CreateNewRound(1)
	fmt.Println("Game starts!")
}

func (g *Game) IsClientInputValid(inputGuess string) bool {

	if input.IsValidGuess(inputGuess, g.CurrentRound.Level.Code.Runes) {
		g.CurrentRound.Level.Trys = g.CurrentRound.Level.Trys + 1
		return true
	} else {
		return false
	}
}

func (g *Game) ClientGuess(clientInput string) GuessResponse {
	if g.IsClientInputValid(clientInput) {

		l := compare.CompareNormalMode(&g.CurrentRound.Level.Code, clientInput)
		fmt.Println(l)

		if compare.CompareRightGuess(&g.CurrentRound.Level.Code, clientInput) {
			evaluatedGuess := compare.CompareEasyMode(&g.CurrentRound.Level.Code, clientInput)
			solved := true
			g.CurrentRound.Level.Endtimer()
			message := "Correct Guess"
			gr := GuessResponse{
				Message:        message,
				Solved:         solved,
				EvaluatedGuess: evaluatedGuess,
				Roundstate:     *g.CurrentRound,
			}
			return gr
		} else {
			evaluatedGuess := compare.CompareEasyMode(&g.CurrentRound.Level.Code, clientInput)
			solved := false
			message := "Try Again"
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
		g.CurrentRound.UpdateCounter()
		g.CurrentRound.Level = level.CreateNewLevel(g.CurrentRound.Counter)
		message := "New Level"
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
		message := "no new level"
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
		fmt.Println("Error serializing the response:", err)
		return nil
	}
	return responseJSON

}
