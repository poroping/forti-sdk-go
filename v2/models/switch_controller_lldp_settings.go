package models

const SwitchControllerLldpSettingsPath = "switch-controller/lldp-settings/"

type SwitchControllerLldpSettings struct {
	DeviceDetection     *string `json:"device-detection,omitempty"`
	FastStartInterval   *int64  `json:"fast-start-interval,omitempty"`
	ManagementInterface *string `json:"management-interface,omitempty"`
	TxHold              *int64  `json:"tx-hold,omitempty"`
	TxInterval          *int64  `json:"tx-interval,omitempty"`
}
