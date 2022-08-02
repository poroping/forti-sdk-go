package models

const UserQuarantinePath = "user/quarantine/"

type UserQuarantine struct {
	FirewallGroups *string                  `json:"firewall-groups,omitempty"`
	Quarantine     *string                  `json:"quarantine,omitempty"`
	Targets        *[]UserQuarantineTargets `json:"targets,omitempty"`
	TrafficPolicy  *string                  `json:"traffic-policy,omitempty"`
}

const UserQuarantineTargetsPath = "user/quarantine/targets/"

type UserQuarantineTargets struct {
	Description *string                      `json:"description,omitempty"`
	Entry       *string                      `json:"entry,omitempty"`
	Macs        *[]UserQuarantineTargetsMacs `json:"macs,omitempty"`
}

const UserQuarantineTargetsMacsPath = "user/quarantine/targets/macs/"

type UserQuarantineTargetsMacs struct {
	Description *string `json:"description,omitempty"`
	Drop        *string `json:"drop,omitempty"`
	Mac         *string `json:"mac,omitempty"`
	Parent      *string `json:"parent,omitempty"`
}
