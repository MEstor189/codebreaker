package main

/* import (
	"github.com/gin-gonic/gin"
) */

import (
	"backend/internal/code"
	"backend/internal/round"
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

	if start {
		c := code.GenerateCode()
		c.CodeString()
	}

}
