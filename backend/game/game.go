package game

import (
	"backend/internal/round"
)

/* type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
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

func HandleInput(){

} */

func Game() {
	start := round.StartRound()

	if start {
		r := round.CreateNewRound(0)
		r.RoundLoop()
	}

}
