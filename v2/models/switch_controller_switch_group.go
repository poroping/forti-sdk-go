package models

const SwitchControllerSwitchGroupPath = "switch-controller/switch-group/"

type SwitchControllerSwitchGroup struct {
	Description *string                               `json:"description,omitempty"`
	Fortilink   *string                               `json:"fortilink,omitempty"`
	Members     *[]SwitchControllerSwitchGroupMembers `json:"members,omitempty"`
	Name        *string                               `json:"name,omitempty"`
}

type SwitchControllerSwitchGroupMembers struct {
	Name     *string `json:"name,omitempty"`
	SwitchId *string `json:"switch-id,omitempty"`
}
