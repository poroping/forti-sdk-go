package models

const SwitchControllerMacPolicyPath = "switch-controller/mac-policy/"

type SwitchControllerMacPolicy struct {
	BouncePortLink *string `json:"bounce-port-link,omitempty"`
	Foscount       *string `json:"foscount,omitempty"`
	Description    *string `json:"description,omitempty"`
	Fortilink      *string `json:"fortilink,omitempty"`
	Name           *string `json:"name,omitempty"`
	TrafficPolicy  *string `json:"traffic-policy,omitempty"`
	Vlan           *string `json:"vlan,omitempty"`
}
