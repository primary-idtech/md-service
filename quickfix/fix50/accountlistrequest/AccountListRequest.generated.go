package accountlistrequest

import (
	"github.com/quickfixgo/quickfix"
	"md-service/quickfix/enum"
	"md-service/quickfix/field"
	"md-service/quickfix/fixt11"
	"md-service/quickfix/tag"
)

// AccountListRequest is the fix50 AccountListRequest type, MsgType = UALR.
type AccountListRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AccountListRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) AccountListRequest {
	return AccountListRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m AccountListRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a AccountListRequest initialized with the required fields for AccountListRequest.
func New(accountrequestid field.AccountRequestIDField, accountlistrequesttype field.AccountListRequestTypeField) (m AccountListRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("UALR"))
	m.Set(accountrequestid)
	m.Set(accountlistrequesttype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg AccountListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "UALR", r
}

// SetAccount sets Account, Tag 1.
func (m AccountListRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m AccountListRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetAccountType sets AccountType, Tag 581.
func (m AccountListRequest) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetAccountRequestID sets AccountRequestID, Tag 7110.
func (m AccountListRequest) SetAccountRequestID(v string) {
	m.Set(field.NewAccountRequestID(v))
}

// SetAccountListRequestType sets AccountListRequestType, Tag 7111.
func (m AccountListRequest) SetAccountListRequestType(v enum.AccountListRequestType) {
	m.Set(field.NewAccountListRequestType(v))
}

// GetAccount gets Account, Tag 1.
func (m AccountListRequest) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m AccountListRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountType gets AccountType, Tag 581.
func (m AccountListRequest) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountRequestID gets AccountRequestID, Tag 7110.
func (m AccountListRequest) GetAccountRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.AccountRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountListRequestType gets AccountListRequestType, Tag 7111.
func (m AccountListRequest) GetAccountListRequestType() (v enum.AccountListRequestType, err quickfix.MessageRejectError) {
	var f field.AccountListRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m AccountListRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m AccountListRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasAccountType returns true if AccountType is present, Tag 581.
func (m AccountListRequest) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasAccountRequestID returns true if AccountRequestID is present, Tag 7110.
func (m AccountListRequest) HasAccountRequestID() bool {
	return m.Has(tag.AccountRequestID)
}

// HasAccountListRequestType returns true if AccountListRequestType is present, Tag 7111.
func (m AccountListRequest) HasAccountListRequestType() bool {
	return m.Has(tag.AccountListRequestType)
}
