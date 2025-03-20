package level

import (
	"backend/internal/code"
	"fmt"
	"time"
)

type Level struct {
	Lvl       int
	Code      code.Code
	Solved    bool
	Trys      int
	StartTime time.Time
	EndTime   time.Duration
}

func CreateNewLevel(lvl int) Level {
	c := code.GenerateCode()
	l := Level{
		Code:      c,
		Lvl:       lvl,
		Solved:    false,
		Trys:      0,
		StartTime: time.Now(),
	}
	fmt.Println("Das aktuelle level ist: ", lvl)
	return l
}
func (l *Level) Endtimer() {
	timer := l.StartTime
	l.EndTime = time.Since(timer)
}
