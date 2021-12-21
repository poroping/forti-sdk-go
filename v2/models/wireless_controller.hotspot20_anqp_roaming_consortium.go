package models

const WirelessControllerhotspot20AnqpRoamingConsortiumPath = "wireless-controller.hotspot20/anqp-roaming-consortium/"

type WirelessControllerhotspot20AnqpRoamingConsortium struct {
	Name   *string                                                   `json:"name,omitempty"`
	OiList *[]WirelessControllerhotspot20AnqpRoamingConsortiumOiList `json:"oi-list,omitempty"`
}

type WirelessControllerhotspot20AnqpRoamingConsortiumOiList struct {
	Comment *string  `json:"comment,omitempty"`
	Index   *float64 `json:"index,omitempty"`
	Oi      *string  `json:"oi,omitempty"`
}
