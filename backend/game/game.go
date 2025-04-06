package game

import (
	"backend/internal/compare"
	"backend/internal/dto"
	"backend/internal/input"
	"backend/internal/level"
	"backend/internal/round"
	"backend/internal/score"
	"encoding/json"
	"fmt"
)

type GuessResponse struct {
	Message string        `json:"message"`
	State   dto.ServerDTO `json:"state"`
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
		if g.CurrentRound.Level.Trys == 1 {
			g.CurrentRound.Level.StartTimer()
		}
		return true
	} else {
		return false
	}
}

func (g *Game) ClientGuess(clientInput string) GuessResponse {
	if g.IsClientInputValid(clientInput) {

		if compare.CompareRightGuess(&g.CurrentRound.Level.Code, clientInput) {
			evaluatedGuess := compare.CompareNormalMode(&g.CurrentRound.Level.Code, clientInput)
			solved := true
			g.CurrentRound.Level.Endtimer()

			fmt.Println(g.CurrentRound.Level.EndTime)
			fmt.Println(g.CurrentRound.Level.Difficulty.Timer)
			g.CurrentRound.Level.LvLScore = score.CalculateScore(g.CurrentRound.Level.Lvl, g.CurrentRound.Level.Trys, g.CurrentRound.Level.EndTime, int(g.CurrentRound.Level.Difficulty.Timer))
			g.CurrentRound.RoundScore += g.CurrentRound.Level.LvLScore

			dto := dto.GenerateDTO(g.CurrentRound, solved, evaluatedGuess)
			message := "Correct Guess"
			gr := GuessResponse{
				Message: message,
				State:   dto,
			}
			return gr
		} else {
			evaluatedGuess := compare.CompareNormalMode(&g.CurrentRound.Level.Code, clientInput)
			solved := false
			dto := dto.GenerateDTO(g.CurrentRound, solved, evaluatedGuess)
			message := "Try Again"
			gr := GuessResponse{
				Message: message,
				State:   dto,
			}
			return gr
		}
	} else {
		dto := dto.GenerateDTO(g.CurrentRound, false, compare.EvaluatedGuessIsEmpty())
		gr := GuessResponse{
			Message: "invalid guess",
			State:   dto,
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
		dto := dto.GenerateDTO(g.CurrentRound, solved, evaluatedGuess)

		gr := GuessResponse{
			Message: message,
			State:   dto,
		}
		return gr
	} else {
		message := "no new level"
		solved := true
		evaluatedGuess := compare.EvaluatedGuessIsEmpty()
		dto := dto.GenerateDTO(g.CurrentRound, solved, evaluatedGuess)

		gr := GuessResponse{
			Message: message,
			State:   dto,
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

func SerializeStartResponse(response *round.Round) []byte {

	dto := dto.GenerateDTO(response, false, compare.EvaluatedGuessIsEmpty())

	responseJSON, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("Error serializing the response:", err)
		return nil
	}
	return responseJSON
}
