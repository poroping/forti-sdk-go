package models

const WirelessControllerSsidPolicyPath = "wireless-controller/ssid-policy/"

type WirelessControllerSsidPolicy struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Vlan        *string `json:"vlan,omitempty"`
}
