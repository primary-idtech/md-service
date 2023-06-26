package fix

import (
	"bytes"
	"strings"
	"testing"

	"md-service/quickfix/fix50/marketdatasnapshotfullrefresh"

	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertFIXMarketData(t *testing.T) {
	txt := "8=FIXT.1.1|9=293|35=W|34=11|49=ROFX|52=20230618-00:31:37.232|56=sfernandez7468|55=DLR/JUN23|207=ROFX|262=DLR/JUN23|264=5|268=6|269=0|270=259.50|271=1|290=0|269=1|270=261.00|271=1|290=0|269=2|270=260.00|271=1|272=20230617|273=20:14:14.957|277=U|288=*|289=*|7201=1|269=B|271=1|269=x|271=1000|269=w|270=260000.00|10=013|"
	txt = strings.ReplaceAll(txt, "|", "\x01")

	msg := quickfix.NewMessage()
	quickfix.ParseMessage(msg, bytes.NewBuffer([]byte(txt)))

	fixMd := marketdatasnapshotfullrefresh.FromMessage(msg)

	converter := &mdConverter{}

	md, err := converter.convert(&fixMd)
	assert.NoError(t, err)
	assert.Equal(t, "DLR/JUN23", md.Symbol)

	// Expected datetime is 20230617 20:14:14.957
	assert.Equal(t, "20230617 20:14:14.957", md.Datetime.Format("20060102 15:04:05.000"))

	// Expected bid is 259.50
	assert.Equal(t, decimal.NewNullDecimal(decimal.RequireFromString("259.50")), md.Bid)

	// Expected ask is 261.00
	assert.Equal(t, decimal.NewNullDecimal(decimal.RequireFromString("261.00")), md.Ask)

	// Expected last is 260.00
	assert.Equal(t, decimal.NewNullDecimal(decimal.RequireFromString("260.00")), md.Last)
}

func TestConvertFIXMarketData_RequiredFieldMissing(t *testing.T) {
	txt := "8=FIXT.1.1|9=134|35=W|34=12|49=ROFX|52=20230626-02:08:42.283|56=sfernandez7468|55=DLR/JUN23|207=ROFX|262=DLR/JUN23|264=5|268=2|269=0|290=0|269=1|290=0|10=161|"
	txt = strings.ReplaceAll(txt, "|", "\x01")

	msg := quickfix.NewMessage()
	quickfix.ParseMessage(msg, bytes.NewBuffer([]byte(txt)))

	fixMd := marketdatasnapshotfullrefresh.FromMessage(msg)
	converter := &mdConverter{}

	md, err := converter.convert(&fixMd)
	assert.NoError(t, err)
	assert.Equal(t, "DLR/JUN23", md.Symbol)
}
