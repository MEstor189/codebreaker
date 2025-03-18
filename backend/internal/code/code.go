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
	//Length int
	//Difficulty int
}

func GenerateCode() Code {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	codeValue := randomizer.Randomize(r, 4)
	return Code{
		Value: codeValue,
		Seed:  seed,
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
