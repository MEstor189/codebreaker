package websocket

import (
	"backend/game"
	"encoding/json"
	"fmt"
	"net/http"

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
	Type    string `json:"type"`
	Content string `json:"content"`
}

func StartWebSocketServer() {
	manager := &Manager{
		Clients:    make(map[*Client]bool),
		HighScores: []int{},
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocketConnection(w, r, manager)
	})

	fmt.Println("WebSocket-Server läuft...")

}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request, manager *Manager) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Fehler beim Upgrade der Verbindung:", err)
		return
	}
	defer conn.Close()

	client := &Client{
		Conn:         conn,
		GameInstance: game.NewGame(),
	}
	manager.Clients[client] = true

	// Kommunikationsloop für den Client
	for {
		// Warte auf eine Nachricht vom Client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Fehler beim Lesen der Nachricht:", err)
			break
		}

		// Rufe die processMessage-Funktion auf, um die eingehende Nachricht zu verarbeiten
		ProcessMessage(client, manager, msg)
	}
}

func ProcessMessage(client *Client, manager *Manager, msg []byte) {
	var incomingMessage Message

	err := json.Unmarshal(msg, &incomingMessage)
	if err != nil {
		fmt.Println("Ungültige Nachricht:", err)
		return
	}

	// Nutze den Type der Nachricht, um zu entscheiden, was zu tun ist
	switch incomingMessage.Type {
	case "start":
		HandleStartGame(client.Conn, client.GameInstance)
		fmt.Println(client.Conn.RemoteAddr().String())

	case "guess":
		HandleGuess(client.Conn, client.GameInstance, incomingMessage.Content)

	case "test":
		HandleNextLevel(client.Conn, client.GameInstance, true)

	default:
		fmt.Println("Unbekannter Nachrichtentyp:", incomingMessage.Type)
	}
}
