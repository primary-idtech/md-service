package pubsub

import (
	"md-service/pkg/fix"
	"md-service/pkg/lvc"
	"md-service/pkg/model"
	"md-service/quickfix/fix50/marketdatasnapshotfullrefresh"
	"testing"
)

func TestPubsub_Subscribe(t *testing.T) {
	lvc := lvc.NewLVC()
	fixMdCh := make(chan *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh)
	app := fix.NewApplication("", "", fixMdCh)
	mdChannel := make(chan *model.MarketData)
	ps := NewPubsub(lvc, app, mdChannel).(*pubsub)

	client1 := &model.Client{ID: "client1", Ch: make(chan *model.MarketData)}
	client2 := &model.Client{ID: "client2", Ch: make(chan *model.MarketData)}

	// Subscribe client1 to symbol1
	symbol1 := "symbol1"
	md1 := ps.Subscribe(client1, symbol1)
	if md1 != nil {
		t.Errorf("Expected nil market data, but got %v", md1)
	}

	// Subscribe client2 to symbol1
	md2 := ps.Subscribe(client2, symbol1)
	if md2 != nil {
		t.Errorf("Expected nil market data, but got %v", md2)
	}

	// Check subscriptions
	clients1, ok1 := ps.subscriptions.Load(symbol1)
	if !ok1 {
		t.Errorf("Expected subscription for symbol1, but not found")
	}
	if len(clients1.([]*model.Client)) != 2 {
		t.Errorf("Expected 2 clients for symbol1, but got %d", len(clients1.([]*model.Client)))
	}
	if clients1.([]*model.Client)[0].ID != "client1" || clients1.([]*model.Client)[1].ID != "client2" {
		t.Errorf("Expected clients [client1, client2] for symbol1, but got %v", clients1)
	}

	// Subscribe client1 to symbol2
	symbol2 := "symbol2"
	md3 := ps.Subscribe(client1, symbol2)
	if md3 != nil {
		t.Errorf("Expected nil market data, but got %v", md3)
	}

	// Check subscriptions
	clients2, ok2 := ps.subscriptions.Load(symbol2)
	if !ok2 {
		t.Errorf("Expected subscription for symbol2, but not found")
	}
	if len(clients2.([]*model.Client)) != 1 {
		t.Errorf("Expected 1 client for symbol2, but got %d", len(clients2.([]*model.Client)))
	}
	if clients2.([]*model.Client)[0].ID != "client1" {
		t.Errorf("Expected client1 for symbol2, but got %v", clients2)
	}
}

func TestPubsub_Unsubscribe(t *testing.T) {
	lvc := lvc.NewLVC()
	fixMdCh := make(chan *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh)
	app := fix.NewApplication("", "", fixMdCh)
	mdChannel := make(chan *model.MarketData)
	ps := NewPubsub(lvc, app, mdChannel).(*pubsub)

	client1 := &model.Client{ID: "client1", Ch: make(chan *model.MarketData)}
	client2 := &model.Client{ID: "client2", Ch: make(chan *model.MarketData)}

	// Subscribe client1 and client2 to symbol1
	symbol1 := "symbol1"
	ps.Subscribe(client1, symbol1)
	ps.Subscribe(client2, symbol1)

	// Unsubscribe client1 from symbol1
	ps.Unsubscribe(client1, symbol1)

	// Check subscriptions
	clients1, ok1 := ps.subscriptions.Load(symbol1)
	if !ok1 {
		t.Errorf("Expected subscription for symbol1, but not found")
	}
	if len(clients1.([]*model.Client)) != 1 {
		t.Errorf("Expected 1 client for symbol1, but got %d", len(clients1.([]*model.Client)))
	}
	if clients1.([]*model.Client)[0].ID != "client2" {
		t.Errorf("Expected client2 for symbol1, but got %v", clients1)
	}

	// Unsubscribe client2 from symbol1
	ps.Unsubscribe(client2, symbol1)

	// Check subscriptions
	_, ok2 := ps.subscriptions.Load(symbol1)
	if ok2 {
		t.Errorf("Expected no subscriptions for symbol1, but found")
	}

	// Unsubscribe client1 from symbol1 again (should not panic)
	ps.Unsubscribe(client1, symbol1)
}

func TestPubsub_Disconnect(t *testing.T) {
	lvc := lvc.NewLVC()
	fixMdCh := make(chan *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh)
	app := fix.NewApplication("", "", fixMdCh)
	mdChannel := make(chan *model.MarketData)
	ps := NewPubsub(lvc, app, mdChannel).(*pubsub)

	client1 := &model.Client{ID: "client1", Ch: make(chan *model.MarketData)}
	client2 := &model.Client{ID: "client2", Ch: make(chan *model.MarketData)}

	// Subscribe client1 and client2 to symbol1
	symbol1 := "symbol1"
	ps.Subscribe(client1, symbol1)
	ps.Subscribe(client2, symbol1)

	// Disconnect client1
	ps.Disconnect(client1)

	// Check clients
	_, ok1 := ps.clients.Load(client1.ID)
	if ok1 {
		t.Errorf("Expected client1 to be disconnected, but found")
	}
	_, ok2 := ps.clients.Load(client2.ID)
	if !ok2 {
		t.Errorf("Expected client2 to be connected, but not found")
	}

	// Check subscriptions
	clients1, ok3 := ps.subscriptions.Load(symbol1)
	if !ok3 {
		t.Errorf("Expected subscription for symbol1, but not found")
	}
	if len(clients1.([]*model.Client)) != 1 {
		t.Errorf("Expected 1 client for symbol1, but got %d", len(clients1.([]*model.Client)))
	}
	if clients1.([]*model.Client)[0].ID != "client2" {
		t.Errorf("Expected client2 for symbol1, but got %v", clients1)
	}

	// Disconnect client2
	ps.Disconnect(client2)

	// Check clients
	_, ok4 := ps.clients.Load(client2.ID)
	if ok4 {
		t.Errorf("Expected client2 to be disconnected, but found")
	}

	// Check subscriptions
	_, ok5 := ps.subscriptions.Load(symbol1)
	if ok5 {
		t.Errorf("Expected no subscriptions for symbol1, but found")
	}
}

func TestPubsub_Publish(t *testing.T) {
	lvc := lvc.NewLVC()
	fixMdCh := make(chan *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh)
	app := fix.NewApplication("", "", fixMdCh)
	mdChannel := make(chan *model.MarketData)
	ps := NewPubsub(lvc, app, mdChannel).(*pubsub)

	client1 := &model.Client{ID: "client1", Ch: make(chan *model.MarketData)}
	client2 := &model.Client{ID: "client2", Ch: make(chan *model.MarketData)}

	// Subscribe client1 and client2 to symbol1
	symbol1 := "symbol1"
	ps.Subscribe(client1, symbol1)
	ps.Subscribe(client2, symbol1)

	// Publish market data for symbol1
	md := &model.MarketData{Symbol: symbol1}
	go ps.Publish(md.Symbol, md)

	// Check client1
	md1 := <-client1.Ch
	if md1 != md {
		t.Errorf("Expected market data %v for client1, but got %v", md, md1)
	}

	// Check client2
	md2 := <-client2.Ch
	if md2 != md {
		t.Errorf("Expected market data %v for client2, but got %v", md, md2)
	}
}
