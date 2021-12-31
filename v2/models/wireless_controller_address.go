package models

const WirelessControllerAddressPath = "wireless-controller/address/"

type WirelessControllerAddress struct {
	Id     *string `json:"id,omitempty"`
	Mac    *string `json:"mac,omitempty"`
	Policy *string `json:"policy,omitempty"`
}
