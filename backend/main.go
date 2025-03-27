package main

import (
	"backend/websocket"
	"fmt"
	"net/http"
)

func main() {

	websocket.StartWebSocketServer()

	port := "8080"
	fmt.Printf("Server runs on http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
