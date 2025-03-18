package main

/* import (
	"github.com/gin-gonic/gin"
) */

import (
	"backend/internal/code"
	"backend/internal/compare"
	"backend/internal/input"
	"backend/internal/output"
	"backend/internal/round"
	"fmt"
)

/* func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
*/

func main() {

	start := round.StartRound()
	running := false
	c := code.GenerateCode()
	p := &c

	if start {
		c.CodeString()
		running = true
	}

	for running {
		fmt.Println("Gib einen Guess ab.")
		userInput := input.GetUserInput()
		if input.IsValidGuess(userInput) {
			if compare.CompareRightGuess(p, userInput) {
				output.OutputRightGuess()
				break
			} else {
				output.OutputWrongGuess(compare.Compare(p, userInput))
			}
		} else {
			fmt.Println("Falscher Input!")
		}
	}

}
