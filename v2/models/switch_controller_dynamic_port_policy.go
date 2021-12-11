package models

const SwitchControllerDynamicPortPolicyPath = "switch-controller/dynamic-port-policy/"

type SwitchControllerDynamicPortPolicy struct {
	Description *string                                   `json:"description,omitempty"`
	Fortilink   *string                                   `json:"fortilink,omitempty"`
	Name        *string                                   `json:"name,omitempty"`
	Policy      []SwitchControllerDynamicPortPolicyPolicy `json:"policy,omitempty"`
}

type SwitchControllerDynamicPortPolicyPolicy struct {
	The8021x       *string                                                `json:"802-1x,omitempty"`
	BouncePortLink *string                                                `json:"bounce-port-link,omitempty"`
	Category       *string                                                `json:"category,omitempty"`
	Description    *string                                                `json:"description,omitempty"`
	Family         *string                                                `json:"family,omitempty"`
	Host           *string                                                `json:"host,omitempty"`
	InterfaceTags  []SwitchControllerDynamicPortPolicyPolicyInterfaceTags `json:"interface-tags,omitempty"`
	LldpProfile    *string                                                `json:"lldp-profile,omitempty"`
	Mac            *string                                                `json:"mac,omitempty"`
	Name           *string                                                `json:"name,omitempty"`
	QosPolicy      *string                                                `json:"qos-policy,omitempty"`
	Status         *string                                                `json:"status,omitempty"`
	Type           *string                                                `json:"type,omitempty"`
	VlanPolicy     *string                                                `json:"vlan-policy,omitempty"`
}

type SwitchControllerDynamicPortPolicyPolicyInterfaceTags struct {
	TagName *string `json:"tag-name,omitempty"`
}
