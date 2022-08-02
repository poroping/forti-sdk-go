package models

const SwitchControllerStpInstancePath = "switch-controller/stp-instance/"

type SwitchControllerStpInstance struct {
	Id        *string                                 `json:"id,omitempty"`
	VlanRange *[]SwitchControllerStpInstanceVlanRange `json:"vlan-range,omitempty"`
}

const SwitchControllerStpInstanceVlanRangePath = "switch-controller/stp-instance/vlan-range/"

type SwitchControllerStpInstanceVlanRange struct {
	VlanName *string `json:"vlan-name,omitempty"`
}
