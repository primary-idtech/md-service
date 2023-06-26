package accountlist

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/quickfix"
	"md-service/quickfix/enum"
	"md-service/quickfix/field"
	"md-service/quickfix/fixt11"
	"md-service/quickfix/tag"
)

// AccountList is the fix50 AccountList type, MsgType = UALT.
type AccountList struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AccountList from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) AccountList {
	return AccountList{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m AccountList) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a AccountList initialized with the required fields for AccountList.
func New(accountrequestid field.AccountRequestIDField, accountrequestresult field.AccountRequestResultField) (m AccountList) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("UALT"))
	m.Set(accountrequestid)
	m.Set(accountrequestresult)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg AccountList, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "UALT", r
}

// SetText sets Text, Tag 58.
func (m AccountList) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetLastFragment sets LastFragment, Tag 893.
func (m AccountList) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

// SetAccountRequestID sets AccountRequestID, Tag 7110.
func (m AccountList) SetAccountRequestID(v string) {
	m.Set(field.NewAccountRequestID(v))
}

// SetAccountRequestResult sets AccountRequestResult, Tag 7112.
func (m AccountList) SetAccountRequestResult(v enum.AccountRequestResult) {
	m.Set(field.NewAccountRequestResult(v))
}

// SetNoRelatedAcc sets NoRelatedAcc, Tag 7113.
func (m AccountList) SetNoRelatedAcc(f NoRelatedAccRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m AccountList) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastFragment gets LastFragment, Tag 893.
func (m AccountList) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountRequestID gets AccountRequestID, Tag 7110.
func (m AccountList) GetAccountRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.AccountRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountRequestResult gets AccountRequestResult, Tag 7112.
func (m AccountList) GetAccountRequestResult() (v enum.AccountRequestResult, err quickfix.MessageRejectError) {
	var f field.AccountRequestResultField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedAcc gets NoRelatedAcc, Tag 7113.
func (m AccountList) GetNoRelatedAcc() (NoRelatedAccRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedAccRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m AccountList) HasText() bool {
	return m.Has(tag.Text)
}

// HasLastFragment returns true if LastFragment is present, Tag 893.
func (m AccountList) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

// HasAccountRequestID returns true if AccountRequestID is present, Tag 7110.
func (m AccountList) HasAccountRequestID() bool {
	return m.Has(tag.AccountRequestID)
}

// HasAccountRequestResult returns true if AccountRequestResult is present, Tag 7112.
func (m AccountList) HasAccountRequestResult() bool {
	return m.Has(tag.AccountRequestResult)
}

// HasNoRelatedAcc returns true if NoRelatedAcc is present, Tag 7113.
func (m AccountList) HasNoRelatedAcc() bool {
	return m.Has(tag.NoRelatedAcc)
}

// NoRelatedAcc is a repeating group element, Tag 7113.
type NoRelatedAcc struct {
	*quickfix.Group
}

// SetAccount sets Account, Tag 1.
func (m NoRelatedAcc) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetAccountName sets AccountName, Tag 7114.
func (m NoRelatedAcc) SetAccountName(v string) {
	m.Set(field.NewAccountName(v))
}

// SetAccountType sets AccountType, Tag 581.
func (m NoRelatedAcc) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetPersonID sets PersonID, Tag 7121.
func (m NoRelatedAcc) SetPersonID(value decimal.Decimal, scale int32) {
	m.Set(field.NewPersonID(value, scale))
}

// SetDealingCapacity sets DealingCapacity, Tag 1048.
func (m NoRelatedAcc) SetDealingCapacity(v string) {
	m.Set(field.NewDealingCapacity(v))
}

// SetAccountRiskCheck sets AccountRiskCheck, Tag 7125.
func (m NoRelatedAcc) SetAccountRiskCheck(v bool) {
	m.Set(field.NewAccountRiskCheck(v))
}

// SetPartyID sets PartyID, Tag 448.
func (m NoRelatedAcc) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetAccountStatus sets AccountStatus, Tag 7126.
func (m NoRelatedAcc) SetAccountStatus(v bool) {
	m.Set(field.NewAccountStatus(v))
}

// SetNoMarketAlias sets NoMarketAlias, Tag 7122.
func (m NoRelatedAcc) SetNoMarketAlias(f NoMarketAliasRepeatingGroup) {
	m.SetGroup(f)
}

// GetAccount gets Account, Tag 1.
func (m NoRelatedAcc) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountName gets AccountName, Tag 7114.
func (m NoRelatedAcc) GetAccountName() (v string, err quickfix.MessageRejectError) {
	var f field.AccountNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountType gets AccountType, Tag 581.
func (m NoRelatedAcc) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPersonID gets PersonID, Tag 7121.
func (m NoRelatedAcc) GetPersonID() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PersonIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDealingCapacity gets DealingCapacity, Tag 1048.
func (m NoRelatedAcc) GetDealingCapacity() (v string, err quickfix.MessageRejectError) {
	var f field.DealingCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountRiskCheck gets AccountRiskCheck, Tag 7125.
func (m NoRelatedAcc) GetAccountRiskCheck() (v bool, err quickfix.MessageRejectError) {
	var f field.AccountRiskCheckField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyID gets PartyID, Tag 448.
func (m NoRelatedAcc) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountStatus gets AccountStatus, Tag 7126.
func (m NoRelatedAcc) GetAccountStatus() (v bool, err quickfix.MessageRejectError) {
	var f field.AccountStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMarketAlias gets NoMarketAlias, Tag 7122.
func (m NoRelatedAcc) GetNoMarketAlias() (NoMarketAliasRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMarketAliasRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasAccount returns true if Account is present, Tag 1.
func (m NoRelatedAcc) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasAccountName returns true if AccountName is present, Tag 7114.
func (m NoRelatedAcc) HasAccountName() bool {
	return m.Has(tag.AccountName)
}

// HasAccountType returns true if AccountType is present, Tag 581.
func (m NoRelatedAcc) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasPersonID returns true if PersonID is present, Tag 7121.
func (m NoRelatedAcc) HasPersonID() bool {
	return m.Has(tag.PersonID)
}

// HasDealingCapacity returns true if DealingCapacity is present, Tag 1048.
func (m NoRelatedAcc) HasDealingCapacity() bool {
	return m.Has(tag.DealingCapacity)
}

// HasAccountRiskCheck returns true if AccountRiskCheck is present, Tag 7125.
func (m NoRelatedAcc) HasAccountRiskCheck() bool {
	return m.Has(tag.AccountRiskCheck)
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoRelatedAcc) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasAccountStatus returns true if AccountStatus is present, Tag 7126.
func (m NoRelatedAcc) HasAccountStatus() bool {
	return m.Has(tag.AccountStatus)
}

// HasNoMarketAlias returns true if NoMarketAlias is present, Tag 7122.
func (m NoRelatedAcc) HasNoMarketAlias() bool {
	return m.Has(tag.NoMarketAlias)
}

// NoMarketAlias is a repeating group element, Tag 7122.
type NoMarketAlias struct {
	*quickfix.Group
}

// SetMarketSegmentID sets MarketSegmentID, Tag 1300.
func (m NoMarketAlias) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

// SetMarketAliasName sets MarketAliasName, Tag 7123.
func (m NoMarketAlias) SetMarketAliasName(v string) {
	m.Set(field.NewMarketAliasName(v))
}

// GetMarketSegmentID gets MarketSegmentID, Tag 1300.
func (m NoMarketAlias) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketAliasName gets MarketAliasName, Tag 7123.
func (m NoMarketAlias) GetMarketAliasName() (v string, err quickfix.MessageRejectError) {
	var f field.MarketAliasNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300.
func (m NoMarketAlias) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

// HasMarketAliasName returns true if MarketAliasName is present, Tag 7123.
func (m NoMarketAlias) HasMarketAliasName() bool {
	return m.Has(tag.MarketAliasName)
}

// NoMarketAliasRepeatingGroup is a repeating group, Tag 7122.
type NoMarketAliasRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMarketAliasRepeatingGroup returns an initialized, NoMarketAliasRepeatingGroup.
func NewNoMarketAliasRepeatingGroup() NoMarketAliasRepeatingGroup {
	return NoMarketAliasRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMarketAlias,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MarketSegmentID), quickfix.GroupElement(tag.MarketAliasName)})}
}

// Add create and append a new NoMarketAlias to this group.
func (m NoMarketAliasRepeatingGroup) Add() NoMarketAlias {
	g := m.RepeatingGroup.Add()
	return NoMarketAlias{g}
}

// Get returns the ith NoMarketAlias in the NoMarketAliasRepeatinGroup.
func (m NoMarketAliasRepeatingGroup) Get(i int) NoMarketAlias {
	return NoMarketAlias{m.RepeatingGroup.Get(i)}
}

// NoRelatedAccRepeatingGroup is a repeating group, Tag 7113.
type NoRelatedAccRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedAccRepeatingGroup returns an initialized, NoRelatedAccRepeatingGroup.
func NewNoRelatedAccRepeatingGroup() NoRelatedAccRepeatingGroup {
	return NoRelatedAccRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedAcc,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AccountName), quickfix.GroupElement(tag.AccountType), quickfix.GroupElement(tag.PersonID), quickfix.GroupElement(tag.DealingCapacity), quickfix.GroupElement(tag.AccountRiskCheck), quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.AccountStatus), NewNoMarketAliasRepeatingGroup()})}
}

// Add create and append a new NoRelatedAcc to this group.
func (m NoRelatedAccRepeatingGroup) Add() NoRelatedAcc {
	g := m.RepeatingGroup.Add()
	return NoRelatedAcc{g}
}

// Get returns the ith NoRelatedAcc in the NoRelatedAccRepeatinGroup.
func (m NoRelatedAccRepeatingGroup) Get(i int) NoRelatedAcc {
	return NoRelatedAcc{m.RepeatingGroup.Get(i)}
}
