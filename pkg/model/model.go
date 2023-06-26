package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type MarketData struct {
	Symbol   string
	Datetime time.Time
	Bid      decimal.NullDecimal
	Ask      decimal.NullDecimal
	Last     decimal.NullDecimal
}

type Client struct {
	ID string
	Ch chan *MarketData
}
