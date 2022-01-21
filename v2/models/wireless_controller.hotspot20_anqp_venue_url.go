package models

const WirelessControllerHotspot20AnqpVenueUrlPath = "wireless-controller.hotspot20/anqp-venue-url/"

type WirelessControllerHotspot20AnqpVenueUrl struct {
	Name      *string                                             `json:"name,omitempty"`
	ValueList *[]WirelessControllerHotspot20AnqpVenueUrlValueList `json:"value-list,omitempty"`
}

type WirelessControllerHotspot20AnqpVenueUrlValueList struct {
	Index  *int64  `json:"index,omitempty"`
	Number *int64  `json:"number,omitempty"`
	Value  *string `json:"value,omitempty"`
}
