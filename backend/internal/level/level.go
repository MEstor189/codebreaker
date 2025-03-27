package level

import (
	"backend/internal/code"
	"backend/internal/difficulty"
	"time"
)

type Level struct {
	Lvl        int
	Code       code.Code
	Solved     bool
	Trys       int
	StartTime  time.Time
	EndTime    time.Duration
	Difficulty difficulty.Difficulty
}

func CreateNewLevel(lvl int) Level {
	difficulty := difficulty.GenerateDifficulty(lvl)

	c := code.GenerateCode(difficulty.CodeLength, difficulty.PSC)
	l := Level{
		Code:       c,
		Lvl:        lvl,
		Solved:     false,
		Trys:       0,
		StartTime:  time.Now(),
		Difficulty: difficulty,
	}
	return l
}
func (l *Level) Endtimer() {
	timer := l.StartTime
	l.EndTime = time.Since(timer)
}
