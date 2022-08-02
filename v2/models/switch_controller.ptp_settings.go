package models

const SwitchControllerPtpSettingsPath = "switch-controller.ptp/settings/"

type SwitchControllerPtpSettings struct {
	Mode *string `json:"mode,omitempty"`
}
