package websocket

import (
	"backend/db"
	"backend/game"
	"context"
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

	responseJSON := game.SerializeStartResponse(gameInstance.CurrentRound)
	err := conn.WriteMessage(websocket.TextMessage, []byte(responseJSON))
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

func HandleHighscoreEntry(conn *websocket.Conn, gameInstance *game.Game, playername string, score int) {
	connDB, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer connDB.Close(context.Background())

	he := db.CreateHighscoreEntry(playername, score)
	err = db.InsertHighscore(connDB, he)
	if err != nil {
		fmt.Errorf("Error inserting: %w", err)
	} else {
		fmt.Println("Highscore etnry inserted!")
	}

}
