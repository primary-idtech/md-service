package main

import (
	"md-service/pkg/fix"
	"os"
	"os/signal"

	"github.com/quickfixgo/quickfix"
)

func main() {
	username := os.Getenv("FIX_USERNAME")
	password := os.Getenv("FIX_PASSWORD")

	app := fix.NewApplication(username, password)

	storeFactory := quickfix.NewMemoryStoreFactory()

	settingsReader, err := os.Open("./quickfix.cfg")
	if err != nil {
		panic(err)
	}
	defer settingsReader.Close()

	appSettings, err := quickfix.ParseSettings(settingsReader)
	if err != nil {
		panic(err)
	}

	logFactory := fix.NewLogFactory()

	initiator, err := quickfix.NewInitiator(app, storeFactory, appSettings, logFactory)
	if err != nil {
		panic(err)
	}

	err = initiator.Start()
	if err != nil {
		panic(err)
	}

	// Start HTTP server
	// TODO

	// Wait for interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	initiator.Stop()
}
