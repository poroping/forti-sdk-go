package models

const SwitchControllerNacDevicePath = "switch-controller/nac-device/"

type SwitchControllerNacDevice struct {
	Description      *string `json:"description,omitempty"`
	Id               *int64  `json:"id,omitempty"`
	LastKnownPort    *string `json:"last-known-port,omitempty"`
	LastKnownSwitch  *string `json:"last-known-switch,omitempty"`
	LastSeen         *int64  `json:"last-seen,omitempty"`
	Mac              *string `json:"mac,omitempty"`
	MacPolicy        *string `json:"mac-policy,omitempty"`
	MatchedNacPolicy *string `json:"matched-nac-policy,omitempty"`
	PortPolicy       *string `json:"port-policy,omitempty"`
	Status           *string `json:"status,omitempty"`
}
