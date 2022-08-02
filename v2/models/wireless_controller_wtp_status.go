package models

const WirelessControllerWtpStatusPath = "wireless-controller/wtp-status/"

type WirelessControllerWtpStatus struct {
	WtpId *string `json:"<wtp-id>,omitempty"`
}
