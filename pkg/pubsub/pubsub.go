package pubsub

import (
	"md-service/pkg/fix"
	"md-service/pkg/lvc"
	"md-service/pkg/model"
	"sync"
)

type Pubsub interface {
	Subscribe(client *model.Client, symbol string) *model.MarketData
	Unsubscribe(client *model.Client, symbol string)
	Disconnect(client *model.Client)
	Publish(symbol string, marketData *model.MarketData)
	Start()
}

type pubsub struct {
	lvc       lvc.LVC
	app       fix.Application
	mdChannel <-chan *model.MarketData

	clients       sync.Map // map[string]*model.Client
	subscriptions sync.Map // map[string][]*model.Client
}

func NewPubsub(lvc lvc.LVC, app fix.Application, mdChannel <-chan *model.MarketData) Pubsub {
	return &pubsub{
		lvc:       lvc,
		app:       app,
		mdChannel: mdChannel,
	}
}

func (ps *pubsub) Subscribe(client *model.Client, symbol string) *model.MarketData {
	// Add client to clients
	ps.clients.Store(client.ID, client)

	// Add client to subscriptions
	clients, loaded := ps.subscriptions.LoadOrStore(symbol, []*model.Client{})
	clients = append(clients.([]*model.Client), client)
	ps.subscriptions.Store(symbol, clients)

	// Subscribe symbol to FIX if not loaded
	if !loaded {
		ps.app.SendMarketDataRequest(symbol)
	}

	// Return market data
	return ps.lvc.GetMarketData(symbol)
}

func (ps *pubsub) Unsubscribe(client *model.Client, symbol string) {
	// Remove client from subscriptions
	clients, _ := ps.subscriptions.LoadOrStore(symbol, []*model.Client{})

	for i, c := range clients.([]*model.Client) {
		if c.ID == client.ID {
			clients = append(clients.([]*model.Client)[:i], clients.([]*model.Client)[i+1:]...)
			break
		}
	}
	ps.subscriptions.Store(symbol, clients)

	// TODO Remove symbol from LVC if no clients are subscribed
	// TODO Unsubscribe symbol from FIX if no clients are subscribed

	// Remove symbol if no clients are subscribed
	if len(clients.([]*model.Client)) == 0 {
		ps.subscriptions.Delete(symbol)
	}
}

func (ps *pubsub) Disconnect(client *model.Client) {
	// Remove client from clients
	ps.clients.Delete(client.ID)

	// Remove client from subscriptions
	ps.subscriptions.Range(func(key, value interface{}) bool {
		clients := value.([]*model.Client)
		for i, c := range clients {
			if c.ID == client.ID {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		ps.subscriptions.Store(key, clients)

		if len(clients) == 0 {
			ps.subscriptions.Delete(key)
		}

		return true
	})
}

func (ps *pubsub) Publish(symbol string, marketData *model.MarketData) {
	// Update LVC
	ps.lvc.UpdateMarketData(marketData)

	// Publish market data to clients
	clients, _ := ps.subscriptions.LoadOrStore(symbol, []*model.Client{})
	for _, client := range clients.([]*model.Client) {
		client.Ch <- marketData
	}
}

func (ps *pubsub) Start() {
	go func() {
		for md := range ps.mdChannel {
			ps.Publish(md.Symbol, md)
		}
	}()
}
