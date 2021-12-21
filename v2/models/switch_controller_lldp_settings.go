package models

const SwitchControllerLldpSettingsPath = "switch-controller/lldp-settings/"

type SwitchControllerLldpSettings struct {
	DeviceDetection     *string  `json:"device-detection,omitempty"`
	FastStartInterval   *float64 `json:"fast-start-interval,omitempty"`
	ManagementInterface *string  `json:"management-interface,omitempty"`
	TxHold              *float64 `json:"tx-hold,omitempty"`
	TxInterval          *float64 `json:"tx-interval,omitempty"`
}
