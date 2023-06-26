package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type MarketData struct {
	Symbol   string              `json:"symbol"`
	Datetime time.Time           `json:"datetime"`
	Bid      decimal.NullDecimal `json:"bid"`
	Ask      decimal.NullDecimal `json:"ask"`
	Last     decimal.NullDecimal `json:"last"`
}

type Client struct {
	ID string
	Ch chan *MarketData
}
