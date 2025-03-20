package main

import (
	"backend/game"
)

/*
	 var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Lasse alle Verbindungen zu (inkl. CORS, eventuell anpassen)
		},
	}
*/
func main() {

	game.Game()

	/* 	// Register WebSocket Route
		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			handleWebSocketConnection(w, r, gameInstance)
		})

		// Starte den HTTP-Server
		port := "8080"
		fmt.Printf("WebSocket-Server läuft auf ws://localhost:%s/ws\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			fmt.Println("Fehler beim Starten des Servers:", err)
		}

	}

	func handleWebSocketConnection(w http.ResponseWriter, r *http.Request, gameInstance *game.Game) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Fehler beim Upgrade der Verbindung:", err)
			return
		}
		defer conn.Close()

		// Hier kannst du Kommunikation mit dem Client behandeln
		// Das könnte z.B. ein Loop sein, um Nachrichten zu empfangen und zu verarbeiten.
		for {
			var msg string
			// Warte auf eine Nachricht vom Client
			err := conn.ReadMessage(&msg)
			if err != nil {
				fmt.Println("Fehler beim Lesen der Nachricht:", err)
				break
			}

			// Hier kannst du die Nachricht verarbeiten und mit der Spiel-Logik kommunizieren
			gameInstance.ProcessMessage(msg)

			// Sende eine Antwort zurück (optional)
			if err := conn.WriteMessage(websocket.TextMessage, []byte("Nachricht empfangen: "+msg)); err != nil {
				fmt.Println("Fehler beim Senden der Nachricht:", err)
				break
			}
		} */
}
