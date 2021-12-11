package models

const WirelessControllerAddressPath = "wireless-controller/address/"

type WirelessControllerAddress struct {
	Fosid  *string `json:"fosid,omitempty"`
	Mac    *string `json:"mac,omitempty"`
	Policy *string `json:"policy,omitempty"`
}
