package instrmtlegexecgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp2/legpreallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties3"
)

//New returns an initialized InstrmtLegExecGrp instance
func New() *InstrmtLegExecGrp {
	var m InstrmtLegExecGrp
	return &m
}

//NoLegs is a repeating group in InstrmtLegExecGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegLastPx is a non-required field for NoLegs.
	LegLastPx *float64 `fix:"637"`
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
	//LegSettlCurrency is a non-required field for NoLegs.
	LegSettlCurrency *string `fix:"675"`
	//LegLastForwardPoints is a non-required field for NoLegs.
	LegLastForwardPoints *float64 `fix:"1073"`
	//LegCalculatedCcyLastQty is a non-required field for NoLegs.
	LegCalculatedCcyLastQty *float64 `fix:"1074"`
	//LegGrossTradeAmt is a non-required field for NoLegs.
	LegGrossTradeAmt *float64 `fix:"1075"`
	//NestedParties3 is a non-required component for NoLegs.
	NestedParties3 *nestedparties3.NestedParties3
	//LegAllocID is a non-required field for NoLegs.
	LegAllocID *string `fix:"1366"`
	//LegPreAllocGrp is a non-required component for NoLegs.
	LegPreAllocGrp *legpreallocgrp.LegPreAllocGrp
	//LegVolatility is a non-required field for NoLegs.
	LegVolatility *float64 `fix:"1379"`
	//LegDividendYield is a non-required field for NoLegs.
	LegDividendYield *float64 `fix:"1381"`
	//LegCurrencyRatio is a non-required field for NoLegs.
	LegCurrencyRatio *float64 `fix:"1383"`
	//LegExecInst is a non-required field for NoLegs.
	LegExecInst *string `fix:"1384"`
	//LegLastQty is a non-required field for NoLegs.
	LegLastQty *float64 `fix:"1418"`
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegQty(v float64)                                  { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetLegPositionEffect(v string)                        { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int)                       { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetLegRefID(v string)                                 { m.LegRefID = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }
func (m *NoLegs) SetLegLastPx(v float64)                               { m.LegLastPx = &v }
func (m *NoLegs) SetLegOrderQty(v float64)                             { m.LegOrderQty = &v }
func (m *NoLegs) SetLegSettlCurrency(v string)                         { m.LegSettlCurrency = &v }
func (m *NoLegs) SetLegLastForwardPoints(v float64)                    { m.LegLastForwardPoints = &v }
func (m *NoLegs) SetLegCalculatedCcyLastQty(v float64)                 { m.LegCalculatedCcyLastQty = &v }
func (m *NoLegs) SetLegGrossTradeAmt(v float64)                        { m.LegGrossTradeAmt = &v }
func (m *NoLegs) SetNestedParties3(v nestedparties3.NestedParties3)    { m.NestedParties3 = &v }
func (m *NoLegs) SetLegAllocID(v string)                               { m.LegAllocID = &v }
func (m *NoLegs) SetLegPreAllocGrp(v legpreallocgrp.LegPreAllocGrp)    { m.LegPreAllocGrp = &v }
func (m *NoLegs) SetLegVolatility(v float64)                           { m.LegVolatility = &v }
func (m *NoLegs) SetLegDividendYield(v float64)                        { m.LegDividendYield = &v }
func (m *NoLegs) SetLegCurrencyRatio(v float64)                        { m.LegCurrencyRatio = &v }
func (m *NoLegs) SetLegExecInst(v string)                              { m.LegExecInst = &v }
func (m *NoLegs) SetLegLastQty(v float64)                              { m.LegLastQty = &v }

//InstrmtLegExecGrp is a fix50sp2 Component
type InstrmtLegExecGrp struct {
	//NoLegs is a non-required field for InstrmtLegExecGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegExecGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
