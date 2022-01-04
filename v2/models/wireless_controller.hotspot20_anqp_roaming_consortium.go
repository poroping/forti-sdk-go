package models

const WirelessControllerHotspot20AnqpRoamingConsortiumPath = "wireless-controller.hotspot20/anqp-roaming-consortium/"

type WirelessControllerHotspot20AnqpRoamingConsortium struct {
	Name   *string                                                   `json:"name,omitempty"`
	OiList *[]WirelessControllerHotspot20AnqpRoamingConsortiumOiList `json:"oi-list,omitempty"`
}

type WirelessControllerHotspot20AnqpRoamingConsortiumOiList struct {
	Comment *string `json:"comment,omitempty"`
	Index   *int64  `json:"index,omitempty"`
	Oi      *string `json:"oi,omitempty"`
}
