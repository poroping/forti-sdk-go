package models

const WirelessControllerhotspot20AnqpVenueNamePath = "wireless-controller/hotspot20/anqp-venue-name/"

type WirelessControllerhotspot20AnqpVenueName struct {
	Name      *string                                              `json:"name,omitempty"`
	ValueList *[]WirelessControllerhotspot20AnqpVenueNameValueList `json:"value-list,omitempty"`
}

type WirelessControllerhotspot20AnqpVenueNameValueList struct {
	Index *int64  `json:"index,omitempty"`
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"value,omitempty"`
}
