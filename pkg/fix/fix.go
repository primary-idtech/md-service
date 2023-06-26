package fix

import (
	"md-service/quickfix/enum"
	"md-service/quickfix/fix50/marketdatasnapshotfullrefresh"

	"github.com/quickfixgo/fixt11/logon"
	"github.com/quickfixgo/quickfix"
)

type Application interface {
	quickfix.Application
	SendMarketDataRequest(symbol string) error
}

// Quickfix Application
type application struct {
	username string
	password string
	fixMdCh  chan<- *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh
}

func NewApplication(username, password string, fixMdCh chan<- *marketdatasnapshotfullrefresh.MarketDataSnapshotFullRefresh) Application {
	return &application{
		username: username,
		password: password,
		fixMdCh:  fixMdCh,
	}
}

func (a *application) OnCreate(sessionID quickfix.SessionID) {
}

func (a *application) OnLogon(sessionID quickfix.SessionID) {
}

func (a *application) OnLogout(sessionID quickfix.SessionID) {
}

func (a *application) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	// inject the user credentials into the outgoing Logon message
	if msg.IsMsgTypeOf(string(enum.MsgType_LOGON)) {
		logonMsg := logon.FromMessage(msg)
		logonMsg.SetUsername(a.username)
		logonMsg.SetPassword(a.password)
	}
}

func (a *application) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error {
	return nil
}

func (a *application) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (a *application) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	if msg.IsMsgTypeOf(string(enum.MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH)) {
		md := marketdatasnapshotfullrefresh.FromMessage(msg)
		a.fixMdCh <- &md
	}

	return nil
}

func (a *application) SendMarketDataRequest(symbol string) error {
	return nil
}
