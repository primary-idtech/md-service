package fix

import (
	"bytes"
	"fmt"

	"github.com/quickfixgo/quickfix"
)

type logFactory struct{}

func NewLogFactory() quickfix.LogFactory {
	return &logFactory{}
}

func (f *logFactory) Create() (quickfix.Log, error) {
	return &log{}, nil
}
func (f *logFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	return &log{sessionID: sessionID}, nil
}

type log struct {
	sessionID quickfix.SessionID
}

func (l *log) OnIncoming(s []byte) {
	fmt.Printf("[%s] fix.incoming: %s\n", l.sessionID.String(), fixMsgToString(s))
}

func (l *log) OnOutgoing(s []byte) {
	if isLogonMsg(s) {
		fmt.Printf("[%s] fix.outgoing: LOGON message\n", l.sessionID.String())
	} else {
		fmt.Printf("[%s] fix.outgoing: %s\n", l.sessionID.String(), fixMsgToString(s))
	}
}

func (l *log) OnEvent(s string) {
	fmt.Printf("[%s] fix.event: %s\n", l.sessionID.String(), s)
}

func (l *log) OnEventf(format string, a ...interface{}) {
	l.OnEvent(fmt.Sprintf(format, a...))
}

// Returns true if the message contains the substring "\x0135=A\x01"
func isLogonMsg(s []byte) bool {
	return bytes.Contains(s, []byte("\x0135=A\x01"))
}

// Returns FIX message with "|" instead of "\x01" as delimiter
func fixMsgToString(s []byte) string {
	return string(bytes.Replace(s, []byte("\x01"), []byte("|"), -1))
}
