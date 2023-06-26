package lvc

import (
	"sync"

	"md-service/pkg/model"
)

type LVC interface {
	UpdateMarketData(md *model.MarketData)
	GetMarketData(symbol string) *model.MarketData
}

type lvc struct {
	marketData sync.Map
}

func NewLVC() LVC {
	return &lvc{
		marketData: sync.Map{},
	}
}

func (l *lvc) UpdateMarketData(md *model.MarketData) {
	l.marketData.Store(md.Symbol, md)
}

func (l *lvc) GetMarketData(symbol string) *model.MarketData {
	md, ok := l.marketData.Load(symbol)
	if !ok {
		return nil
	}
	return md.(*model.MarketData)
}
