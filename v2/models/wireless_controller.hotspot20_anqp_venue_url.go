package models

const WirelessControllerhotspot20AnqpVenueUrlPath = "wireless-controller.hotspot20/anqp-venue-url/"

type WirelessControllerhotspot20AnqpVenueUrl struct {
	Name      *string                                             `json:"name,omitempty"`
	ValueList *[]WirelessControllerhotspot20AnqpVenueUrlValueList `json:"value-list,omitempty"`
}

type WirelessControllerhotspot20AnqpVenueUrlValueList struct {
	Index  *float64 `json:"index,omitempty"`
	Number *float64 `json:"number,omitempty"`
	Value  *string  `json:"value,omitempty"`
}
