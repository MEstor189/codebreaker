package round

import (
	"backend/internal/level"
)

type Round struct {
	Level      level.Level
	Counter    int
	Running    bool
	Score      int64
	RoundScore int64
}

func CreateNewRound(counter int) *Round {
	level := level.CreateNewLevel(counter)
	r := Round{
		Level:      level,
		Counter:    counter,
		Running:    true,
		RoundScore: 0,
	}
	return &r
}

func (r *Round) UpdateCounter() {
	r.Counter = r.Counter + 1
}
