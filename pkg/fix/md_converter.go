package fix

import (
	"fmt"
	"md-service/pkg/model"
	"time"

	"md-service/quickfix/fix50/marketdatasnapshotfullrefresh"

	"md-service/quickfix/enum"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type MdConverter interface {
	Start()
}

type mdConverter struct {
	fixMdCh <-chan *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh
	mdCh    chan<- *model.MarketData
}

func NewMdConverter(
	fixMdCh <-chan *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh,
	mdCh chan<- *model.MarketData,
) MdConverter {
	return &mdConverter{
		fixMdCh: fixMdCh,
		mdCh:    mdCh,
	}
}

func (mdc *mdConverter) Start() {
	go func() {
		for fixMd := range mdc.fixMdCh {
			// Convert FIX MarketData to internal MarketData
			md, err := mdc.convert(fixMd)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Publish internal MarketData
			mdc.mdCh <- &md
		}
	}()
}

func (mdc *mdConverter) convert(fixMd *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh) (md model.MarketData, err error) {
	// Symbol
	symbol, err := fixMd.GetSymbol()
	if err != nil {
		return md, err
	}

	md.Symbol = symbol

	entries, err := fixMd.GetNoMDEntries()
	if err != nil {
		return md, err
	}

	for i := 0; i < entries.Len(); i++ {
		var err error
		entry := entries.Get(i)

		entryType, err := entry.GetMDEntryType()
		if err != nil {
			return md, err
		}

		if entryType == enum.MDEntryType_TRADE {
			var err error

			// Trade price
			price, err := entry.GetMDEntryPx()
			if err != nil {
				return md, err
			}

			md.Last = decimal.NullDecimal{
				Decimal: price,
				Valid:   true,
			}

			datetime, err := mdc.extractDateTimeFromEntry(entry)
			if err != nil {
				return md, err
			}

			md.Datetime = datetime
		}

		if entryType == enum.MDEntryType_BID {
			// Bid price
			price, err := entry.GetMDEntryPx()
			if err != nil {
				return md, err
			}

			md.Bid = decimal.NullDecimal{
				Decimal: price,
				Valid:   true,
			}
		}

		if entryType == enum.MDEntryType_OFFER {
			// Ask price
			price, err := entry.GetMDEntryPx()
			if err != nil {
				return md, err
			}

			md.Ask = decimal.NullDecimal{
				Decimal: price,
				Valid:   true,
			}
		}
	}

	return md, nil
}

func (mdc *mdConverter) extractDateTimeFromEntry(entry marketdatasnapshotfullrefresh.NoMDEntries) (time.Time, error) {
	datetime := time.Time{}

	mdDate, err := entry.GetMDEntryDate()
	if err != nil {
		return datetime, errors.Wrap(err, "failed to extract date from entry")
	}

	mdTime, err := entry.GetMDEntryTime()
	if err != nil {
		return datetime, errors.Wrap(err, "failed to extract time from entry")
	}

	return time.Parse("20060102-15:04:05", fmt.Sprintf("%s-%s", mdDate, mdTime))
}
