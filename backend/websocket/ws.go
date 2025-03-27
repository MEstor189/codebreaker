package websocket

import (
	"backend/game"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn         *websocket.Conn
	GameInstance *game.Game
}

type Manager struct {
	Clients    map[*Client]bool
	HighScores []int
}

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type GuessContent struct {
	PressedSymbols []string `json:"pressedSymbols"`
}

func StartWebSocketServer() {
	manager := &Manager{
		Clients:    make(map[*Client]bool),
		HighScores: []int{},
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocketConnection(w, r, manager)
	})

	fmt.Println("WebSocket-Server is running...")

}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request, manager *Manager) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	client := &Client{
		Conn:         conn,
		GameInstance: game.NewGame(),
	}
	manager.Clients[client] = true

	for {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error loadung message:", err)
			break
		}
		fmt.Println("Message:", string(msg))

		ProcessMessage(client, manager, msg)
	}
}

func ProcessMessage(client *Client, manager *Manager, msg []byte) {
	var incomingMessage Message

	err := json.Unmarshal(msg, &incomingMessage)
	if err != nil {
		fmt.Println("Invalid Message:", err)
		return
	}

	switch incomingMessage.Type {

	case "start":
		HandleStartRound(client.Conn, client.GameInstance)
		fmt.Println(client.Conn.RemoteAddr().String())

	case "guess":
		var guessContent GuessContent
		err := json.Unmarshal(incomingMessage.Payload, &guessContent)
		if err != nil {
			fmt.Println("Error parsing 'guess'-data:", err)
			return
		}
		symbolsString := strings.Join(guessContent.PressedSymbols, "")
		HandleGuess(client.Conn, client.GameInstance, symbolsString)

	case "nextLevel":
		HandleNextLevel(client.Conn, client.GameInstance, true)

	default:
		fmt.Println("Unknown Messagetype:", incomingMessage.Type)
	}
}
