package models

const WirelessControllerApStatusPath = "wireless-controller/ap-status/"

type WirelessControllerApStatus struct {
	Bssid  *string `json:"bssid,omitempty"`
	Id     *int64  `json:"id,omitempty"`
	Ssid   *string `json:"ssid,omitempty"`
	Status *string `json:"status,omitempty"`
}
