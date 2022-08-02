package models

const WirelessControllerHotspot20AnqpVenueNamePath = "wireless-controller.hotspot20/anqp-venue-name/"

type WirelessControllerHotspot20AnqpVenueName struct {
	Name      *string                                              `json:"name,omitempty"`
	ValueList *[]WirelessControllerHotspot20AnqpVenueNameValueList `json:"value-list,omitempty"`
}

const WirelessControllerHotspot20AnqpVenueNameValueListPath = "wireless-controller.hotspot20/anqp-venue-name/value-list/"

type WirelessControllerHotspot20AnqpVenueNameValueList struct {
	Index *int64  `json:"index,omitempty"`
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"value,omitempty"`
}
