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
	marketData map[string]*model.MarketData
	mutex      sync.RWMutex
}

func NewLVC() LVC {
	return &lvc{
		marketData: make(map[string]*model.MarketData),
	}
}

func (l *lvc) UpdateMarketData(md *model.MarketData) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.marketData[md.Symbol] = md
}

func (l *lvc) GetMarketData(symbol string) *model.MarketData {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.marketData[symbol]
}
