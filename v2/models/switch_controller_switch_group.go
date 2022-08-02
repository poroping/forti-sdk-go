package models

const SwitchControllerSwitchGroupPath = "switch-controller/switch-group/"

type SwitchControllerSwitchGroup struct {
	Description *string                               `json:"description,omitempty"`
	Fortilink   *string                               `json:"fortilink,omitempty"`
	Members     *[]SwitchControllerSwitchGroupMembers `json:"members,omitempty"`
	Name        *string                               `json:"name,omitempty"`
}

const SwitchControllerSwitchGroupMembersPath = "switch-controller/switch-group/members/"

type SwitchControllerSwitchGroupMembers struct {
	SwitchId *string `json:"switch-id,omitempty"`
}
