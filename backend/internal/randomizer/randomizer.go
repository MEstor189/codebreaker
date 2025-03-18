package randomizer

import (
	"bytes"
	"math/rand"
)

var latin = []rune{'1', '2', '3', '4'}

func Randomize(rand *rand.Rand, size int) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString(string(latin[rand.Intn(len(latin))]))
	}
	s := buffer.String()
	return s
}

func GetLatin() []rune {
	return latin
}
