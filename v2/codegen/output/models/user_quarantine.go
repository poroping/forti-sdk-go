package models

const UserQuarantinePath = "user/quarantine/"

type UserQuarantine struct {
	FirewallGroups *string                 `json:"firewall-groups,omitempty"`
	Quarantine     *string                 `json:"quarantine,omitempty"`
	Targets        []UserQuarantineTargets `json:"targets,omitempty"`
	TrafficPolicy  *string                 `json:"traffic-policy,omitempty"`
}

type UserQuarantineTargets struct {
	Description *string                     `json:"description,omitempty"`
	Entry       *string                     `json:"entry,omitempty"`
	Macs        []UserQuarantineTargetsMacs `json:"macs,omitempty"`
}

type UserQuarantineTargetsMacs struct {
	Description *string `json:"description,omitempty"`
	Drop        *string `json:"drop,omitempty"`
	Mac         *string `json:"mac,omitempty"`
	Parent      *string `json:"parent,omitempty"`
}
