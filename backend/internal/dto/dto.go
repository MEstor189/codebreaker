package dto

import (
	"backend/internal/compare"
	"backend/internal/difficulty"
	"backend/internal/round"
)

type ServerDTO struct {
	Solved                 bool
	ComparisonResultNormal compare.ComparisonResultNormal
	Difficulty             difficulty.Difficulty
	Trys                   int
	Runes                  []rune
	Level                  int
	LvLScore               int64
	RoundScore             int64
}

func GenerateDTO(round *round.Round, solved bool, evaluatedGuess compare.ComparisonResultNormal) ServerDTO {
	return ServerDTO{
		Solved:                 solved,
		ComparisonResultNormal: evaluatedGuess,
		Difficulty:             round.Level.Difficulty,
		Trys:                   round.Level.Trys,
		Runes:                  round.Level.Code.Runes,
		Level:                  round.Level.Lvl,
		LvLScore:               round.Level.LvLScore,
		RoundScore:             round.RoundScore,
	}

}
