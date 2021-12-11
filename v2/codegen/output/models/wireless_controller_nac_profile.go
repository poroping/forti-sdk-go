package models

const WirelessControllerNacProfilePath = "wireless-controller/nac-profile/"

type WirelessControllerNacProfile struct {
	Comment        *string `json:"comment,omitempty"`
	Name           *string `json:"name,omitempty"`
	OnboardingVlan *string `json:"onboarding-vlan,omitempty"`
}
