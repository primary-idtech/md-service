package main

import (
	"context"
	"fmt"
	"log"
	"md-service/pkg/fix"
	"md-service/pkg/websocket"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/quickfixgo/quickfix"
)

func main() {
	username := os.Getenv("FIX_USERNAME")
	password := os.Getenv("FIX_PASSWORD")

	app := fix.NewApplication(username, password)

	storeFactory := quickfix.NewMemoryStoreFactory()

	settingsReader, err := os.Open("./quickfix.cfg")
	if err != nil {
		log.Fatal(err)
	}
	defer settingsReader.Close()

	appSettings, err := quickfix.ParseSettings(settingsReader)
	if err != nil {
		log.Fatal(err)
	}

	logFactory := fix.NewLogFactory()

	initiator, err := quickfix.NewInitiator(app, storeFactory, appSettings, logFactory)
	if err != nil {
		log.Fatal(err)
	}

	err = initiator.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Start websocket server
	webSocketHandler := websocket.NewWebSocketHandler()
	http.HandleFunc("/ws", webSocketHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	initiator.Stop()

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
