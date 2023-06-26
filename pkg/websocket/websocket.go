package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"md-service/pkg/model"
	"md-service/pkg/pubsub"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	MdTypeSubscribe   = "md.sub"
	MdTypeUnsubscribe = "md.unsub"
)

type clientMessage struct {
	Type    string   `json:"type"`
	Symbols []string `json:"symbols"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewWebSocketHandler(ps pubsub.Pubsub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		// Create new client
		client := &model.Client{
			ID: uuid.New().String(),
			Ch: make(chan *model.MarketData),
		}

		ctx := r.Context()

		// Start a goroutine to read messages from the channel and send them to WebSocket
		go func() {
			for {
				select {
				case <-ctx.Done():
					// Unsubscribe from all symbols
					ps.Disconnect(client)
					return
				case md := <-client.Ch:
					err = conn.WriteJSON(md)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}
		}()

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

			// Parse message
			clientMessage := clientMessage{}
			err = json.Unmarshal(message, &clientMessage)
			if err != nil {
				log.Println(err)
				return
			}

			switch clientMessage.Type {
			case MdTypeSubscribe:
				// Subscribe to market data
				for _, symbol := range clientMessage.Symbols {
					md := ps.Subscribe(client, symbol)
					if md != nil {
						// Send market data to WebSocket
						err = conn.WriteJSON(md)
						if err != nil {
							log.Println(err)
							return
						}
					}
				}

			case MdTypeUnsubscribe:
				// Unsubscribe from market data
				for _, symbol := range clientMessage.Symbols {
					ps.Unsubscribe(client, symbol)
				}
			}
		}
	}
}
