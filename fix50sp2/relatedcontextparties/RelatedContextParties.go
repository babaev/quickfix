package relatedcontextparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedcontextptyssubgrp"
)

//New returns an initialized RelatedContextParties instance
func New() *RelatedContextParties {
	var m RelatedContextParties
	return &m
}

//NoRelatedContextPartyIDs is a repeating group in RelatedContextParties
type NoRelatedContextPartyIDs struct {
	//RelatedContextPartyID is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyID *string `fix:"1576"`
	//RelatedContextPartyIDSource is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyIDSource *string `fix:"1577"`
	//RelatedContextPartyRole is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyRole *int `fix:"1578"`
	//RelatedContextPtysSubGrp is a non-required component for NoRelatedContextPartyIDs.
	RelatedContextPtysSubGrp *relatedcontextptyssubgrp.RelatedContextPtysSubGrp
}

//NewNoRelatedContextPartyIDs returns an initialized NoRelatedContextPartyIDs instance
func NewNoRelatedContextPartyIDs() *NoRelatedContextPartyIDs {
	var m NoRelatedContextPartyIDs
	return &m
}

func (m *NoRelatedContextPartyIDs) SetRelatedContextPartyID(v string) { m.RelatedContextPartyID = &v }
func (m *NoRelatedContextPartyIDs) SetRelatedContextPartyIDSource(v string) {
	m.RelatedContextPartyIDSource = &v
}
func (m *NoRelatedContextPartyIDs) SetRelatedContextPartyRole(v int) { m.RelatedContextPartyRole = &v }
func (m *NoRelatedContextPartyIDs) SetRelatedContextPtysSubGrp(v relatedcontextptyssubgrp.RelatedContextPtysSubGrp) {
	m.RelatedContextPtysSubGrp = &v
}

//RelatedContextParties is a fix50sp2 Component
type RelatedContextParties struct {
	//NoRelatedContextPartyIDs is a non-required field for RelatedContextParties.
	NoRelatedContextPartyIDs []NoRelatedContextPartyIDs `fix:"1575,omitempty"`
}

func (m *RelatedContextParties) SetNoRelatedContextPartyIDs(v []NoRelatedContextPartyIDs) {
	m.NoRelatedContextPartyIDs = v
}
