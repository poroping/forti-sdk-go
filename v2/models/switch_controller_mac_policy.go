package models

const SwitchControllerMacPolicyPath = "switch-controller/mac-policy/"

type SwitchControllerMacPolicy struct {
	BouncePortLink *string `json:"bounce-port-link,omitempty"`
	Count          *string `json:"count,omitempty"`
	Description    *string `json:"description,omitempty"`
	Fortilink      *string `json:"fortilink,omitempty"`
	Name           *string `json:"name,omitempty"`
	TrafficPolicy  *string `json:"traffic-policy,omitempty"`
	Vlan           *string `json:"vlan,omitempty"`
}
