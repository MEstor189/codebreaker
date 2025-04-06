package main

import (
	"backend/db"
	"backend/websocket"
	"fmt"
	"log"
	"net/http"
)

func main() {

	if err := db.InitDB(); err != nil {
		log.Fatalf("DB init error: %v", err)
	}
	fmt.Println("DB connected!")

	websocket.StartWebSocketServer()

	port := "8080"
	fmt.Printf("Server runs on http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
