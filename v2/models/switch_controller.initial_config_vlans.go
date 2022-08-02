package models

const SwitchControllerInitialConfigVlansPath = "switch-controller.initial-config/vlans/"

type SwitchControllerInitialConfigVlans struct {
	DefaultVlan *string `json:"default-vlan,omitempty"`
	Nac         *string `json:"nac,omitempty"`
	NacSegment  *string `json:"nac-segment,omitempty"`
	Quarantine  *string `json:"quarantine,omitempty"`
	Rspan       *string `json:"rspan,omitempty"`
	Video       *string `json:"video,omitempty"`
	Voice       *string `json:"voice,omitempty"`
}
