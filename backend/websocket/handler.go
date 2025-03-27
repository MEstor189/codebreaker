package websocket

import (
	"backend/game"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type StartResponse struct {
	Message string     `json:"message"`
	Game    *game.Game `json:"game"`
}

func HandleNewGame(conn *websocket.Conn) {
	g := game.NewGame()
	sr := StartResponse{
		Message: "created",
		Game:    g,
	}

	responseJSON, err := json.Marshal(sr)
	if err != nil {
		fmt.Println("Error serializing the message:", err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte(responseJSON))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func HandleStartRound(conn *websocket.Conn, gameInstance *game.Game) {
	gameInstance.StartGame()
	responseJSON, err := json.Marshal(gameInstance.CurrentRound)
	if err != nil {
		fmt.Println("Error serializing the response:", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(responseJSON))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func HandleGuess(conn *websocket.Conn, gameInstance *game.Game, guess string) {

	responseJson := game.SerializeGuessResponse(gameInstance.ClientGuess(guess))
	err := conn.WriteMessage(websocket.TextMessage, []byte(responseJson))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func HandleNextLevel(conn *websocket.Conn, gameInstance *game.Game, nextLevel bool) {

	responseJson := game.SerializeGuessResponse(gameInstance.NextLevel(nextLevel))
	err := conn.WriteMessage(websocket.TextMessage, []byte(responseJson))
	if err != nil {
		fmt.Println("Error sending message", err)
	}
}
