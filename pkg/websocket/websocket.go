package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewWebSocketHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		// Handle WebSocket messages
		for {
			// Read message from WebSocket
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			// Print message to console
			fmt.Printf("Received message: %s\n", message)

			// Write message back to WebSocket
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
