package models

const WirelessControllerVapGroupPath = "wireless-controller/vap-group/"

type WirelessControllerVapGroup struct {
	Comment *string                           `json:"comment,omitempty"`
	Name    *string                           `json:"name,omitempty"`
	Vaps    *[]WirelessControllerVapGroupVaps `json:"vaps,omitempty"`
}

const WirelessControllerVapGroupVapsPath = "wireless-controller/vap-group/vaps/"

type WirelessControllerVapGroupVaps struct {
	Name *string `json:"name,omitempty"`
}
