package difficulty

import (
	"fmt"
	"math"
)

type Difficulty struct {
	CodeLength int
	Timer      float64
	PSC        int
}

func GenerateDifficulty(lvl int) Difficulty {
	return Difficulty{
		CodeLength: CalculateCodeLenght(lvl),
		Timer:      CalculateTimer(lvl),
		PSC:        CalculatePossibleSymbolsCount(lvl),
	}
}

func CalculateTimer(level int) float64 {
	baseTimer := 600.0
	reductionFactor := 122.4

	if level < 5 {
		return baseTimer
	}

	timer := baseTimer - (reductionFactor * math.Log(float64(level)))

	if timer < 10 {
		return 10
	}
	return timer
}

func CalculateCodeLenght(level int) int {
	baseCodeLength := 4
	return baseCodeLength + (level-1)/10
}

func CalculatePossibleSymbolsCount(level int) int {
	basePSC := 4
	fmt.Println("PSC", (basePSC + (level-1)/5))
	return basePSC + (level-1)/5

}

//to do:

//font im frontend ändern für mehr Symbole

//timer nochmal überarbeiten, nimmt am anfang zu stark ab
