//Package collateralreport msg type = BA.
package collateralreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CollateralReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BA"`
	fixt11.Header
	//CollRptID is a required field for CollateralReport.
	CollRptID string `fix:"908"`
	//CollInquiryID is a non-required field for CollateralReport.
	CollInquiryID *string `fix:"909"`
	//CollStatus is a required field for CollateralReport.
	CollStatus int `fix:"910"`
	//TotNumReports is a non-required field for CollateralReport.
	TotNumReports *int `fix:"911"`
	//LastRptRequested is a non-required field for CollateralReport.
	LastRptRequested *bool `fix:"912"`
	//Parties is a non-required component for CollateralReport.
	Parties *parties.Parties
	//Account is a non-required field for CollateralReport.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralReport.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralReport.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralReport.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralReport.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralReport.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp is a non-required component for CollateralReport.
	ExecCollGrp *execcollgrp.ExecCollGrp
	//TrdCollGrp is a non-required component for CollateralReport.
	TrdCollGrp *trdcollgrp.TrdCollGrp
	//Instrument is a non-required component for CollateralReport.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralReport.
	FinancingDetails *financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralReport.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralReport.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralReport.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralReport.
	Currency *string `fix:"15"`
	//InstrmtLegGrp is a non-required component for CollateralReport.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//UndInstrmtGrp is a non-required component for CollateralReport.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//MarginExcess is a non-required field for CollateralReport.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralReport.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralReport.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps is a non-required component for CollateralReport.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
	//Side is a non-required field for CollateralReport.
	Side *string `fix:"54"`
	//MiscFeesGrp is a non-required component for CollateralReport.
	MiscFeesGrp *miscfeesgrp.MiscFeesGrp
	//Price is a non-required field for CollateralReport.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralReport.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralReport.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralReport.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralReport.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralReport.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData is a non-required component for CollateralReport.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations is a non-required component for CollateralReport.
	Stipulations *stipulations.Stipulations
	//SettlInstructionsData is a non-required component for CollateralReport.
	SettlInstructionsData *settlinstructionsdata.SettlInstructionsData
	//TradingSessionID is a non-required field for CollateralReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralReport.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralReport.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralReport.
	ClearingBusinessDate *string `fix:"715"`
	//Text is a non-required field for CollateralReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralReport.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for CollateralReport.
	TransactTime *time.Time `fix:"60"`
	//CollApplType is a non-required field for CollateralReport.
	CollApplType *int `fix:"1043"`
	//FinancialStatus is a non-required field for CollateralReport.
	FinancialStatus *string `fix:"291"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized CollateralReport instance
func New(collrptid string, collstatus int) *Message {
	var m Message
	m.SetCollRptID(collrptid)
	m.SetCollStatus(collstatus)
	return &m
}

func (m *Message) SetCollRptID(v string)                                   { m.CollRptID = v }
func (m *Message) SetCollInquiryID(v string)                               { m.CollInquiryID = &v }
func (m *Message) SetCollStatus(v int)                                     { m.CollStatus = v }
func (m *Message) SetTotNumReports(v int)                                  { m.TotNumReports = &v }
func (m *Message) SetLastRptRequested(v bool)                              { m.LastRptRequested = &v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = &v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                                     { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                                     { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)                            { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                            { m.SecondaryClOrdID = &v }
func (m *Message) SetExecCollGrp(v execcollgrp.ExecCollGrp)                { m.ExecCollGrp = &v }
func (m *Message) SetTrdCollGrp(v trdcollgrp.TrdCollGrp)                   { m.TrdCollGrp = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = &v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)                                   { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)          { m.InstrmtLegGrp = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)          { m.UndInstrmtGrp = &v }
func (m *Message) SetMarginExcess(v float64)                               { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)                              { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)                            { m.CashOutstanding = &v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetMiscFeesGrp(v miscfeesgrp.MiscFeesGrp)                { m.MiscFeesGrp = &v }
func (m *Message) SetPrice(v float64)                                      { m.Price = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetAccruedInterestAmt(v float64)                         { m.AccruedInterestAmt = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64)                      { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)                                  { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)                                    { m.EndCash = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetStipulations(v stipulations.Stipulations) { m.Stipulations = &v }
func (m *Message) SetSettlInstructionsData(v settlinstructionsdata.SettlInstructionsData) {
	m.SettlInstructionsData = &v
}
func (m *Message) SetTradingSessionID(v string)     { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)  { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)          { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)       { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = &v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }
func (m *Message) SetTransactTime(v time.Time)      { m.TransactTime = &v }
func (m *Message) SetCollApplType(v int)            { m.CollApplType = &v }
func (m *Message) SetFinancialStatus(v string)      { m.FinancialStatus = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50, "BA", r
}
