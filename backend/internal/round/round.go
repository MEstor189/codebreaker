package round

import (
	"backend/internal/level"
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

func (r *Round) UpdateCounter() {
	r.Counter = r.Counter + 1
}
