package websocket

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"md-service/pkg/model"
)

type mockPubsub struct {
	mock.Mock
	client *model.Client
}

func (m *mockPubsub) Subscribe(client *model.Client, symbol string) *model.MarketData {
	args := m.Called(client, symbol)
	m.client = client
	return args.Get(0).(*model.MarketData)
}

func (m *mockPubsub) Unsubscribe(client *model.Client, symbol string) {
	m.Called(client, symbol)
}

func (m *mockPubsub) Disconnect(client *model.Client) {
	m.Called(client)
}

func (m *mockPubsub) Publish(symbol string, marketData *model.MarketData) {
	m.Called(symbol, marketData)
	m.client.Ch <- marketData
}

func (m *mockPubsub) Start() {
	m.Called()
}

func TestNewWebSocketHandler(t *testing.T) {
	// Create a new pubsub instance
	ps := &mockPubsub{}

	// Create a test server
	ts := httptest.NewServer(NewWebSocketHandler(ps))
	defer ts.Close()

	// Create a new WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial("ws"+ts.URL[4:], nil)
	assert.NoError(t, err)

	ps.On("Subscribe", mock.Anything, "AAPL").Return(&model.MarketData{
		Symbol: "AAPL",
		Last:   decimal.NewNullDecimal(decimal.RequireFromString("100.0")),
	}).Once()
	ps.On("Subscribe", mock.Anything, "GOOG").Return(&model.MarketData{
		Symbol: "GOOG",
		Last:   decimal.NewNullDecimal(decimal.RequireFromString("200.0")),
	}).Once()

	err = conn.WriteJSON(clientMessage{Type: MdTypeSubscribe, Symbols: []string{"AAPL", "GOOG"}})
	assert.NoError(t, err)

	// Test receiving market data
	md := &model.MarketData{}
	err = conn.ReadJSON(md)
	assert.NoError(t, err)
	assert.Equal(t, "AAPL", md.Symbol)
	assert.Equal(t, "100", md.Last.Decimal.String())

	md = &model.MarketData{}
	err = conn.ReadJSON(md)
	assert.NoError(t, err)
	assert.Equal(t, "GOOG", md.Symbol)
	assert.Equal(t, "200", md.Last.Decimal.String())

	// Publish new market data
	ps.On("Publish", "AAPL", mock.Anything).Once()
	go ps.Publish("AAPL", &model.MarketData{
		Symbol: "AAPL",
		Last:   decimal.NewNullDecimal(decimal.RequireFromString("300.0")),
	})

	// Test receiving published market data
	md = &model.MarketData{}
	err = conn.ReadJSON(md)
	assert.NoError(t, err)
	assert.Equal(t, "AAPL", md.Symbol)
	assert.Equal(t, "300", md.Last.Decimal.String())

	done := make(chan struct{})
	ps.On("Unsubscribe", mock.Anything, "AAPL").Run(func(mock.Arguments) {
		done <- struct{}{}
	}).Once()
	ps.On("Unsubscribe", mock.Anything, "GOOG").Run(func(mock.Arguments) {
		done <- struct{}{}
	}).Once()
	err = conn.WriteJSON(clientMessage{Type: MdTypeUnsubscribe, Symbols: []string{"AAPL", "GOOG"}})
	assert.NoError(t, err)

	<-done
	<-done

	// Test disconnecting from the server
	ps.On("Disconnect", mock.Anything).Run(func(mock.Arguments) {
		done <- struct{}{}
	}).Once()
	conn.Close()

	// Wait for the server to finish processing the messages
	<-done

	// Assert expectations
	ps.AssertExpectations(t)
}
