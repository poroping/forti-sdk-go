package models

const SwitchControllerPortPolicyPath = "switch-controller/port-policy/"

type SwitchControllerPortPolicy struct {
	The8021X       *string `json:"802-1x,omitempty"`
	BouncePortLink *string `json:"bounce-port-link,omitempty"`
	Description    *string `json:"description,omitempty"`
	Fortilink      *string `json:"fortilink,omitempty"`
	LldpProfile    *string `json:"lldp-profile,omitempty"`
	Name           *string `json:"name,omitempty"`
	QosPolicy      *string `json:"qos-policy,omitempty"`
	VlanPolicy     *string `json:"vlan-policy,omitempty"`
}
