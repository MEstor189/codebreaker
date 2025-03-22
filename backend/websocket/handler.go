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
		Message: "erstellt",
		Game:    g,
	}

	responseJSON, err := json.Marshal(sr)
	if err != nil {
		fmt.Println("Fehler beim Serialisieren der Antwort:", err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte(responseJSON))
	if err != nil {
		fmt.Println("Fehler beim Senden der Nachricht:", err)
	}
}

func HandleStartGame(conn *websocket.Conn, gameInstance *game.Game) {
	gameInstance.StartGame()
	err := conn.WriteMessage(websocket.TextMessage, []byte("Spiel gestartet"))
	if err != nil {
		fmt.Println("Fehler beim Senden der Nachricht:", err)
	}
}

func HandleGuess(conn *websocket.Conn, gameInstance *game.Game, guess string) {

	responseJson := game.SerializeGuessResponse(gameInstance.ClientGuess(guess))
	err := conn.WriteMessage(websocket.TextMessage, []byte(responseJson))
	if err != nil {
		fmt.Println("Fehler beim Senden der Nachricht:", err)
	}
}

func HandleNextLevel(conn *websocket.Conn, gameInstance *game.Game, nextLevel bool) {

	responseJson := game.SerializeGuessResponse(gameInstance.NextLevel(nextLevel))
	err := conn.WriteMessage(websocket.TextMessage, []byte(responseJson))
	if err != nil {
		fmt.Println("Fehler beim Senden der Nachricht:", err)
	}
}
