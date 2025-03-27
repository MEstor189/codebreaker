package code

import (
	"backend/internal/randomizer"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Code struct {
	Value string
	Seed  int64
	Runes []rune
}

func GenerateCode(length int, psc int) Code {
	symbolsCollection, err := randomizer.LoadSymbols("symbols.toml")
	if err != nil {
		fmt.Println("Error loading symbols:", err)
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	runes := randomizer.GenerateRandomRunes(symbolsCollection, psc)

	codeValue := randomizer.Randomize(r, length, runes)
	return Code{
		Value: codeValue,
		Seed:  seed,
		Runes: runes,
	}
}

func (c *Code) CodeString() {
	fmt.Println(c.Value)
	fmt.Println(c.Seed)
}

func (c *Code) CodeValueAsSlice() []string {
	codeValueSlice := strings.Split(c.Value, "")
	return codeValueSlice
}
