package models

const SwitchControllerNacSettingsPath = "switch-controller/nac-settings/"

type SwitchControllerNacSettings struct {
	AutoAuth       *string `json:"auto-auth,omitempty"`
	BounceNacPort  *string `json:"bounce-nac-port,omitempty"`
	InactiveTimer  *int64  `json:"inactive-timer,omitempty"`
	LinkDownFlush  *string `json:"link-down-flush,omitempty"`
	Mode           *string `json:"mode,omitempty"`
	Name           *string `json:"name,omitempty"`
	OnboardingVlan *string `json:"onboarding-vlan,omitempty"`
}
