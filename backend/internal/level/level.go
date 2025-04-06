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
	EndTime    int
	Difficulty difficulty.Difficulty
	LvLScore   int64
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
		LvLScore:   0,
	}
	return l
}
func (l *Level) Endtimer() {
	timer := l.StartTime
	l.EndTime = int(time.Since(timer).Seconds())
}

func (l *Level) StartTimer() {
	l.StartTime = time.Now()
}
