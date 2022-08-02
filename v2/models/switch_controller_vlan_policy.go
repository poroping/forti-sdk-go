package models

const SwitchControllerVlanPolicyPath = "switch-controller/vlan-policy/"

type SwitchControllerVlanPolicy struct {
	AllowedVlans    *[]SwitchControllerVlanPolicyAllowedVlans  `json:"allowed-vlans,omitempty"`
	AllowedVlansAll *string                                    `json:"allowed-vlans-all,omitempty"`
	Description     *string                                    `json:"description,omitempty"`
	DiscardMode     *string                                    `json:"discard-mode,omitempty"`
	Fortilink       *string                                    `json:"fortilink,omitempty"`
	Name            *string                                    `json:"name,omitempty"`
	UntaggedVlans   *[]SwitchControllerVlanPolicyUntaggedVlans `json:"untagged-vlans,omitempty"`
	Vlan            *string                                    `json:"vlan,omitempty"`
}

const SwitchControllerVlanPolicyAllowedVlansPath = "switch-controller/vlan-policy/allowed-vlans/"

type SwitchControllerVlanPolicyAllowedVlans struct {
	VlanName *string `json:"vlan-name,omitempty"`
}

const SwitchControllerVlanPolicyUntaggedVlansPath = "switch-controller/vlan-policy/untagged-vlans/"

type SwitchControllerVlanPolicyUntaggedVlans struct {
	VlanName *string `json:"vlan-name,omitempty"`
}
