package models

const WirelessControllerApStatusPath = "wireless-controller/ap-status/"

type WirelessControllerApStatus struct {
	Bssid  *string `json:"bssid,omitempty"`
	Fosid  *int64  `json:"fosid,omitempty"`
	Ssid   *string `json:"ssid,omitempty"`
	Status *string `json:"status,omitempty"`
}
